from presidio_analyzer import AnalyzerEngine
from bs4 import BeautifulSoup
import re

analyzer = AnalyzerEngine()

def sanitize_text(text: str) -> str:
    if not text:
        return text

    # 1. Strip HTML
    text = BeautifulSoup(text, "html.parser").get_text()

    # 2. Run Presidio
    results = analyzer.analyze(text=text, language='en', entities=["PERSON", "EMAIL_ADDRESS"])

    # 3. Safe words (e.g., technical terms, not names)
    safe_words = {
        "access", "test", "hang", "load", "retry", "fix", "devtools",
        "monday", "friday", "chrome", "incognito", "bundle", "patch",
        "tools", "reports", "queue"
    }

    # 4. Filter out safe words and false positives
    redactions = []
    for r in results:
        entity_text = text[r.start:r.end]
        words = entity_text.strip().lower().split()

        # Allow only if all words are in the safe list
        if r.entity_type == "PERSON":
            if all(word in safe_words or len(word) < 3 for word in words):
                continue

        redactions.append((r.start, r.end))

    # 5. Redact Presidio-detected values
    for start, end in sorted(redactions, reverse=True):
        text = text[:start] + "(REDACTED)" + text[end:]

    # 6. Regex-based cleanup for email addresses
    text = re.sub(r'\b[\w\.-]+@[\w\.-]+\.\w+\b', '(REDACTED)', text)
    text = re.sub(r'@\w+\b', '@(REDACTED)', text)

    # 7. Redact From/To/Subject lines
    # Matches: From: Nia <nia@example.com> or From: Kenny Lowe
    text = re.sub(r'(?i)^(From|To|Subject):.*$', r'\1: (REDACTED)', text, flags=re.MULTILINE)

    # 8. Redact names in narrative email phrases (e.g., "Email from Eva to Aju")
    text = re.sub(r'(?i)\b(Email\s+from)\s+([A-Z][a-z]+)\s+(to)\s+([A-Z][a-z]+)\b',
                  r'\1 (REDACTED) \3 (REDACTED)', text)

    return text

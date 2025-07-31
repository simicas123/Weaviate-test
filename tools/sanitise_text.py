from presidio_analyzer import AnalyzerEngine
from presidio_anonymizer import AnonymizerEngine
from bs4 import BeautifulSoup

# Initialise the Presidio Analyzer and Anonymizer
analyzer = AnalyzerEngine()
anonymizer = AnonymizerEngine()

def sanitize_text(text: str) -> str:
    """
    Sanitises the provided text by anonymising any detected personal information.
    
    :param text: The text to be sanitised (description or internal comments).
    :return: Sanitised text with personal information removed/anonymised.
    """
    if not text:
        return text
    
    # Remove HTML from the data
    textNoHtml = BeautifulSoup(text, "html.parser").get_text()
    
    # Analyse the text for personal data
    results = analyzer.analyze(textNoHtml, language='en')

    # Anonymise the detected personal data
    sanitized_text = anonymizer.anonymize(text=textNoHtml, analyzer_results=results).text
    
    return sanitized_text

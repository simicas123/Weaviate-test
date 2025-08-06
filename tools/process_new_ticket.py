from sanitise_text import sanitize_text
from embed_text import embed
from upload_to_weaviate import upload_ticket
import re
import base64
import requests

def extract_image_urls(text):
    if not isinstance(text, str):
        return []
    return re.findall(r'(https?://\S+\.(?:png|jpg|jpeg|gif|bmp)[^\s]*)', text, re.IGNORECASE)

def download_and_encode_image(url):
    try:
        print(f"[Downloading image from]: {url}")
        response = requests.get(url, timeout=10)
        if response.status_code == 200:
            encoded = base64.b64encode(response.content).decode('utf-8')
            print(f"[✓] Encoded image base64 length: {len(encoded)}")
            return encoded
        else:
            print(f"[✗] Failed to download image. Status: {response.status_code}")
    except Exception as e:
        print(f"[✗] Exception during image download: {e}")
    return None

def process_new_ticket(raw_ticket):
    # Extract required fields
    title = sanitize_text(raw_ticket['fields'].get('System.Title', 'N/A'))
    description = sanitize_text(raw_ticket['fields'].get('System.Description', 'N/A'))
    internal_comments = sanitize_text(raw_ticket['fields'].get('Workpro.InternalComments', 'N/A'))
    investigation_outcome = sanitize_text(raw_ticket['fields'].get('Workpro.InvestigationOutcome', 'N/A'))
    root_cause = sanitize_text(raw_ticket['fields'].get('Workpro.RootCause', 'N/A'))
    root_cause_reason = sanitize_text(raw_ticket['fields'].get('Workpro.RootCauseReason', 'N/A'))
    how_fixed = sanitize_text(raw_ticket['fields'].get('Workpro.HowFixed', 'N/A'))
    response_due_date = sanitize_text(raw_ticket['fields'].get('Workpro.ResponseDueDate', 'N/A'))

    full_text = f"{title} - {description} - {internal_comments} - {investigation_outcome} - {root_cause} - {root_cause_reason} - {how_fixed} - {response_due_date}"
    embedding = embed(full_text)

    # Handle images
    image_urls = extract_image_urls(description) + extract_image_urls(internal_comments)
    print(f"[Image URLs] Found: {image_urls}")
    images = [img for url in image_urls if (img := download_and_encode_image(url))]

    print(f"[Ticket Images] {raw_ticket['id']} -> {len(images)} image(s) encoded.")

    ticket = {
        "work_item_id": raw_ticket['id'],
        "title": title,
        "text": full_text,
        "embedding": embedding,
        "images": images
    }

    upload_ticket(ticket)
    print(f"[✓] Uploaded ticket {ticket['work_item_id']}")

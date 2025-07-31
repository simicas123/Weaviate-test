from sanitise_text import sanitize_text
from embed_text import embed
from upload_to_weaviate import upload_ticket

def process_new_ticket(raw_ticket):
    # Extract required fields dynamically
    title = raw_ticket['fields'].get('System.Title', 'N/A')
    description = raw_ticket['fields'].get('System.Description', 'N/A')
    internal_comments = raw_ticket['fields'].get('Workpro.InternalComments', 'N/A')
    investigation_outcome = raw_ticket['fields'].get('Workpro.InvestigationOutcome', 'N/A')
    root_cause = raw_ticket['fields'].get('Workpro.RootCause', 'N/A')
    root_cause_reason = raw_ticket['fields'].get('Workpro.RootCauseReason', 'N/A')
    how_fixed = raw_ticket['fields'].get('Workpro.HowFixed', 'N/A')
    response_due_date = raw_ticket['fields'].get('Workpro.ResponseDueDate', 'N/A')

    full_text = f"{title} - {description} - {internal_comments} - {investigation_outcome} - {root_cause} - {root_cause_reason} - {how_fixed} - {response_due_date}"
    
    clean_text = sanitize_text(full_text)
    embedding = embed(clean_text)

    ticket = {
        "work_item_id": raw_ticket['id'],
        "title": title,
        "text": clean_text,
        "embedding": embedding
    }

    upload_ticket(ticket)
    print(f"[✓] Uploaded ticket {ticket['work_item_id']}")


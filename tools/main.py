from process_and_upload import process
from weaviate_schema import weaviate_schema_process
from scrub import scrub
from fetch_new_ticket import fetch_new_ticket
from sanitise_text import sanitize_text
from embed_text import embed
from get_similar_ticket import get_similar_ticket
# MAIN SCRIPT FOR RUNNING PROGRAM

if __name__ == "__main__":
    # !-------- OPTION TO ASK USER TO GET ALL TICKETS FROM TFS ---------! #
    processTicketsAgain = input("Retrieve all old TFS tickets and process into Weaviate? (Y/N)")
    if processTicketsAgain == "Y":
        #run code for retrieving the tickets
        print(f"Retrieving and processing old TFS tickets...")
        #process_and_upload, run this first
        process()

        #weaviate_schema, run after process and upload
        weaviate_schema_process()

    else:
        print(f"Skipping 'Process Tickets Again', proceeding to next option...")
    
    deleteAllTickets = input("Delete all TFS tickets stored in Weaviate? (Y/N)")
    if deleteAllTickets == "Y":
        print(f"Deleting all stored tickets in Weaviate...")
        #scrub.py, run to scrub tickets
        scrub()
        quit()
    else:
        print(f"Skipping 'Delete All Tickets', proceeding to next option...")

# --------- ASK USER NEW ID WORK_ITEM_ID ---------- #
    print(f"PLEASE INPUT TICKET...")
    work_item_id = input("Enter work_item_id to find similar tickets: ")
 
    # Fetch the raw ticket from TFS
    raw_ticket = fetch_new_ticket(work_item_id)
    if not raw_ticket:
        print("[✗] Could not fetch the ticket.")
        exit()
 
    # Prepare and embed the ticket ||text
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
    # Find similar tickets
    similar_tickets = get_similar_ticket(embedding, work_item_id)
 
    # Print results
    print(f"\n[✓] Top similar tickets to {work_item_id}:")
    for idx, ticket in enumerate(similar_tickets, 1):
        print(f"\n→ Match #{idx}")
        print(f"  ID: {ticket['work_item_id']}")
        print(f"  Title: {ticket['title']}")
        print(f"  Text Preview: {ticket['text'][:200]}...")
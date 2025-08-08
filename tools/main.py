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
 
    # Prepare and embed the ticket (THIS IS FOR THE NEW TICKET)
    title = raw_ticket['fields'].get('System.Title', 'N/A')
    description = raw_ticket['fields'].get('System.Description', 'N/A')
    internal_comments = raw_ticket['fields'].get('Workpro.InternalComments', 'N/A')
    investigation_outcome = raw_ticket['fields'].get('Workpro.InvestigationOutcome', 'N/A')
    root_cause = raw_ticket['fields'].get('Workpro.RootCause', 'N/A')
    root_cause_reason = raw_ticket['fields'].get('Workpro.RootCauseReason', 'N/A')
    how_fixed = raw_ticket['fields'].get('Workpro.HowFixed', 'N/A')
    response_due_date = raw_ticket['fields'].get('Workpro.ResponseDueDate', 'N/A')
    area = raw_ticket['fields'].get('System.AreaPath', 'N/A')

    # Emphasise less or more important fields
    important_fields = f"""
    [Important] Area: {area}
    [Important] Root Cause: {root_cause}
    [Important] Root Cause Reason: {root_cause_reason}
    [Important] How Fixed: {how_fixed}
    [Important] Description: {description}
    [Important] Title: {title}
    """

    less_important_fields = f"""
    [LessImportant] Internal Comments: {internal_comments}
    [LessImportant] Investigation Outcome: {investigation_outcome}
    [LessImportant] Response Due Date: {response_due_date}
    """

    full_text = f"{important_fields}\n{important_fields}\n{less_important_fields}"
    
    clean_text = sanitize_text(full_text)
    embedding = embed(clean_text) # EMBEDDING

    # Ask user how many similar tickets they want returned
    num_tickets = 5        
    while True:
        try:
            num_tickets = int(input("How many similar tickets to retrieve? (5, 10, or 20): "))
            if num_tickets in [5, 10, 20]:
                break
            else:
                print("Please enter a valid number: 5, 10, or 20.")
        except ValueError:
            print("Invalid input. Please enter a number.")


    # Find similar tickets
    similar_tickets = get_similar_ticket(embedding, work_item_id, top_k=num_tickets)
 
    # Print results
    print(f"\n[✓] Top similar tickets to {work_item_id}:")
    for idx, ticket in enumerate(similar_tickets, 1):
        print(f"\n→ Match #{idx}")
        print(f"  ID: {ticket['work_item_id']}")
        print(f"  Title: {ticket['title']}")
        print(f"  Text Preview: {ticket['text']}...")
        print(f"  Similarity Score: {ticket['similarity_score']:.4f}")
        if ticket['images']:
            print(f"  Number of images: {len(ticket['images'])}")
            for idx, img in enumerate(ticket['images']):
                print(f"Image #{idx + 1} Base64 snippet: {img[:60]}...") # Limit image print to 60 chars
        else:
            print("No images available for this ticket.")
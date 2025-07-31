from fetch_work_items import fetch_work_items

from sanitise_text import sanitize_text

from embed_text import embed

from upload_to_weaviate import upload_ticket

def process():

    items = fetch_work_items()

    for item in items:
        print(f"looking at an item!", item["work_item_id"])

        title = item["title"]

        desc = item["description"]

        internalCmts = item["internal_comments"]

        rootCause = item["root_cause"]

        investigationOutcome = item["investigation_outcome"]
        
        rootCauseReason = item["root_cause_reason"]        

        howFixed = item["how_fixed"]

        responseDueDate = item["response_due_date"]
        
        full_text = f"{title} - {desc} - {internalCmts} - {rootCause} - {rootCauseReason} - {investigationOutcome} - {howFixed} - {responseDueDate}"
 
        clean_text = sanitize_text(full_text)

        embedding = embed(clean_text)
 
        ticket = {

            "work_item_id": item["work_item_id"],

            "title": title,

            "text": clean_text,

            "embedding": embedding

        }

        upload_ticket(ticket)

        print(f"Processed ticket {item['work_item_id']}")
 
if __name__ == "__main__":
    print(f"Going to run process().")
    process()
    print(f"Finished processing.")

 
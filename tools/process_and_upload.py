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

        iteration = item["iteration"]

        area = item["area"]

        images = item.get("images", [])

        # Sanitise before concatenation to prevent errors

        title = sanitize_text(title)
        desc = sanitize_text(desc)
        internalCmts = sanitize_text(internalCmts)
        rootCause = sanitize_text(rootCause)
        howFixed = sanitize_text(howFixed)
        rootCauseReason = sanitize_text(rootCauseReason)
        investigationOutcome = sanitize_text(investigationOutcome)
        responseDueDate = sanitize_text(responseDueDate)
        # No need to sanitise area and iteration

        # Make some fields more or less important
        important_fields = f"""
        [Important] Iteration: {iteration}
        [Important] Area: {area}
        [Important] Description: {desc}
        [Important] Root Cause: {rootCause}
        [Important] Root Cause Reason: {rootCauseReason}
        [Important] How Fixed: {howFixed}
        """

        less_important_fields = f"""
        [Important] Title: {title}
        [LessImportant] Internal Comments: {internalCmts}
        [LessImportant] Investigation Outcome: {investigationOutcome}
        [LessImportant] Response Due Date: {responseDueDate}
        """
        
        full_text = f"{important_fields}\n{important_fields}\n{less_important_fields}"
 
        embedding = embed(full_text) # Embed based on important fields, less important fields
 
        ticket = {

            "work_item_id": item["work_item_id"],

            "title": title,

            "text": full_text,

            "embedding": embedding,

            "images": images

        }

        upload_ticket(ticket)

        print(f"Processed ticket {item['work_item_id']}")
 
if __name__ == "__main__":
    print(f"Going to run process().")
    process()
    print(f"Finished processing.")

 
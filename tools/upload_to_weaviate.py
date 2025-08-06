import weaviate
from weaviate.collections.classes.config import Configure
 
# Connect to the local Weaviate instance running on Docker

client = weaviate.connect_to_local()

def upload_ticket(ticket):
    try:
        ticket_collection = client.collections.get("Ticket")
        
        properties_to_insert = {
            "work_item_id": str(ticket["work_item_id"]),
            "title": ticket["title"],
            "text": ticket["text"],
            "images": ticket.get("images", [])
        }

        print(f"[DEBUG] Uploading ticket with properties keys: {list(properties_to_insert.keys())}")
        print(f"[DEBUG] Number of images: {len(properties_to_insert['images'])}")

        # Optionally, print first 1-2 base64 strings truncated for sanity
        if properties_to_insert['images']:
            print(f"[DEBUG] First image base64 snippet: {properties_to_insert['images'][0][:50]}...")

        ticket_collection.data.insert(
            properties=properties_to_insert,
            vector=ticket["embedding"]
        )

        print(f"[✓] Ticket {ticket['work_item_id']} uploaded to Weaviate.")
    except Exception as e:
        print(f"[✗] Error uploading ticket {ticket['work_item_id']}: {e}")

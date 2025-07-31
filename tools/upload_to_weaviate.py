import weaviate
from weaviate.collections.classes.config import Configure
 
# Connect to the local Weaviate instance running on Docker

client = weaviate.connect_to_local()

def upload_ticket(ticket):
    ticket_collection = client.collections.get("Ticket")

    ticket_collection.data.insert(
        properties={
            "work_item_id": str(ticket["work_item_id"]),
            "title": ticket["title"],
            "text": ticket["text"],
        },
        vector=ticket["embedding"]  
    )
import weaviate
import weaviate.classes.config as wc
from process_and_upload import process

from upload_to_weaviate import client

def weaviate_schema_process():
    print(client.is_ready()) #check if client is ready

    existing_collections = client.collections.list_all()
    print("Existing collections:", existing_collections)

    if "Ticket" not in existing_collections:
        print("Creating 'Ticket' collection...")
        ticket_schema = client.collections.create(
            name="Ticket",
            description="Support ticket with ID, title, and text",
            properties=[
                wc.Property(name="work_item_id", data_type=wc.DataType.TEXT, skip_vectorization=True),
                wc.Property(name="title", data_type=wc.DataType.TEXT),
                wc.Property(name="text", data_type=wc.DataType.TEXT),
            ],
            vectorizer_config=None,
        )
    else:
        print("Collection 'Ticket' already exists")

    # Fetch work items from your function
    #work_items = process().items # get from embed text + work id
    collection = client.collections.get("Ticket")
    results = collection.query.fetch_objects()
    work_items = [obj.properties for obj in results.objects]

    # Insert fetched work items into the existing collection "Ticket"
    #for item in work_items:
    #    client.collections["Ticket"].objects.create(
    #        data_object={
    #            "work_item_id": item["work_item_id"],
    #            "title": item["title"],
    #            "text": item["text"]
    #        }
    #    )

    print(f"Retrieved {len(work_items)} tickets from Weaviate")
    for item in work_items:
        print(item)

    print(f"Inserted tickets into the 'Ticket' collection.")
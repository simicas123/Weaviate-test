from upload_to_weaviate import client

# Get collection and results
ticket_collection = client.collections.get("Ticket")
results = ticket_collection.query.fetch_objects(limit=5, include_vector=True)

# Script used to query and return tickets we have processed
print(f"Total objects found: {len(results.objects)}\n")

for obj in results.objects:
    props = obj.properties
    work_item_id = props.get("work_item_id", "Unknown")
    images = props.get("images")

    print(f"\nWork Item ID: {work_item_id}")
    print(f"Raw images field (type={type(images)}):\n{images}")
    break  # Just inspect the first object

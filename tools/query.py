from upload_to_weaviate import client
 
# Get the collection
ticket_collection = client.collections.get("Ticket")
 
# Fetch up to 100 objects
results = ticket_collection.query.fetch_objects(limit=100, include_vector=True)
 
# Print each ticket grouped by work_item_id
for obj in results.objects:
    work_item_id = obj.properties.get("work_item_id", "Unknown ID")
    print(f"Work Item ID: {work_item_id}")
    print("Properties:")
    for key, value in obj.properties.items():
        print(f"  {key}: {value}")
    print("Embedding vector:")
    if obj.vector:
        print(f"  {obj.vector}...") 
    else:
        print("  None (no embedding)")
    print("-" * 60)
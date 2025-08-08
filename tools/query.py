from upload_to_weaviate import client

# Get collection and results
ticket_collection = client.collections.get("Ticket")
results = ticket_collection.query.fetch_objects(limit=500, include_vector=True)

print(f"Total objects found: {len(results.objects)}\n")

# Print each ticket's fields
for idx, obj in enumerate(results.objects, 1):
    props = obj.properties
    print(f"\n--- Ticket #{idx} ---")

    for key, value in props.items():
        if key == "images":
            if isinstance(value, list):
                image_preview = [img[:10] + "..." for img in value]
                print(f"{key} (list of {len(value)} images): {image_preview}")
            else:
                print(f"{key}: {value}")  # Not a list? Just print as-is
        else:
            print(f"{key}: {value}")

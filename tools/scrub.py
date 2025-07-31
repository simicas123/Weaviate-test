from upload_to_weaviate import client

# Delete all objects by applying a filter that always matches
def scrub():
    client.collections.delete(
        "Ticket"
    )
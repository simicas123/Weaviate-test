from upload_to_weaviate import client

def get_similar_ticket(embedding, input_id, top_k=5):
    """
    Finds tickets most similar to a given vector embedding using Weaviate Collections API (v4.16.4),
    excluding the ticket with the input_id from the results.
    
    Args:
        embedding (list): The vector embedding of the input ticket.
        input_id (str or int): The ID of the ticket to exclude from results.
        top_k (int): Number of similar tickets to retrieve (before filtering).
    """
    collection = client.collections.get("Ticket")

    # Query for more than needed in case we filter one or more out
    raw_results = collection.query.near_vector(
        near_vector=embedding,
        limit=top_k + 5  # Fetch more to compensate for filtering
    )

    similar_tickets = []
    for obj in raw_results.objects:
        work_item_id = obj.properties.get("work_item_id")

        # Skip the ticket if it matches the input ID
        if str(work_item_id) == str(input_id):
            continue

        similar_tickets.append({
            "work_item_id": work_item_id,
            "title": obj.properties.get("title"),
            "text": obj.properties.get("text"),
            # "distance": obj.distance  # Not available in this API
        })

        # Stop once we've gathered top_k valid results
        if len(similar_tickets) == top_k:
            break

    return similar_tickets

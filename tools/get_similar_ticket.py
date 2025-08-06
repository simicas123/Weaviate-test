from upload_to_weaviate import client
# top_k=5
def get_similar_ticket(embedding, input_id, top_k):
    """
    Finds tickets most similar to a given vector embedding using Weaviate Collections API (v4.16.4),
    excluding the ticket with the input_id from the results.
    
    Args:
        embedding (list): The vector embedding of the input ticket.
        input_id (str or int): The ID of the ticket to exclude from results.
        top_k (int): Number of similar tickets to retrieve (before filtering).
    """
    collection = client.collections.get("Ticket")

    raw_results = collection.query.near_vector(
        near_vector=embedding,
        limit=top_k + 5,
        return_metadata=["distance"]
    )

    print(f"[DEBUG] Raw query result contains {len(raw_results.objects)} objects")

    similar_tickets = []
    for obj in raw_results.objects:
        work_item_id = obj.properties.get("work_item_id")

        if str(work_item_id) == str(input_id):
            continue

        distance = getattr(obj.metadata, "distance", 1.0)
        similarity_score = 1 - distance

        images = obj.properties.get("images", [])
        print(f"[DEBUG] Ticket ID {work_item_id} has {len(images)} images")

        similar_tickets.append({
            "work_item_id": work_item_id,
            "title": obj.properties.get("title"),
            "text": obj.properties.get("text"),
            "images": images,
            "similarity_score": similarity_score
        })

        if len(similar_tickets) == top_k:
            break

    return similar_tickets

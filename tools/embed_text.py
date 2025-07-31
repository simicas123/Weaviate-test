from sentence_transformers import SentenceTransformer
import os
os.environ["NO_PROXY"] = "localhost,127.0.0.1"

model = SentenceTransformer('all-MiniLM-L6-v2')
 
def embed(text):
    return model.encode(text).tolist()
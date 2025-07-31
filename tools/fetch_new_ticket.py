import requests
import os
from requests.auth import HTTPBasicAuth

# Configuration
TFS_URL_BASE = "https://tfs-2018.casltd.com/tfs/Workpro/Workpro/_apis/wit/workitems"
PAT = "v7smcijbmq4syjzihngjzw3nacvtuoca7xhrnwmxkbrnyzouaf5q"
auth = HTTPBasicAuth('', PAT)
os.environ["NO_PROXY"] = "localhost,127.0.0.1"

def fetch_new_ticket(work_item_id):
    """Fetch a single ticket by ID from TFS."""
    url = f"{TFS_URL_BASE}/{work_item_id}?api-version=3.0"
    response = requests.get(url, auth=auth)

    if response.status_code == 200:
        return response.json()
    else:
        print(f"[✗] Failed to fetch work item {work_item_id}. Status: {response.status_code}")
        return None

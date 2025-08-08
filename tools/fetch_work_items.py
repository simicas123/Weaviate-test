import os
import requests
from requests.auth import HTTPBasicAuth
from datetime import datetime, timedelta
from sanitise_text import sanitize_text  # Import the sanitise function
import time
from requests.exceptions import RequestException, ConnectionError
import base64
import re
import html

os.environ["NO_PROXY"] = "localhost,127.0.0.1"

# TFS Organisation URL
tfs_url = "https://tfs-2018.casltd.com/tfs/Workpro/Workpro/_apis/wit/wiql?api-version=3.0"

# Personal Access Token 
PAT = "v7smcijbmq4syjzihngjzw3nacvtuoca7xhrnwmxkbrnyzouaf5q"
auth = HTTPBasicAuth('', PAT)

# Calculate the date 14 days ago
date_14_days_ago = (datetime.now() - timedelta(days=14)).strftime('%Y-%m-%d')

# WIQL query to fetch closed work items from the last 14 days in the Support Queue
wiql_query = {
    "query": f"""
    SELECT [System.Id], [System.Title], [System.State], [System.ChangedDate], [System.IterationPath]
    FROM WorkItems
    WHERE [System.State] = 'Closed'
    AND [System.ChangedDate] >= '{date_14_days_ago}'
    AND [System.IterationPath] = 'Workpro\\Helpdesk\\Support Queue'
    ORDER BY [System.ChangedDate] DESC
    """
}

def request_with_retries(method, url, retries=2, delay=2, **kwargs):
    attempt = 0
    while attempt <= retries:
        try:
            response = requests.request(method, url, **kwargs)
            response.raise_for_status()
            return response
        except (RequestException, ConnectionError) as e:
            print(f"[RETRYING. Retrying: {attempt+1}/{retries+1}] {method.upper()} {url} failed with error: {e}")
            attempt += 1
            if attempt > retries:
                print("Max retries reached. Connection error, is server down?")
                raise
            time.sleep(delay)

all_tickets = [] # list to hold tickets

# Helper to extract image URL
def extract_all_image_urls(text):
    # Match all <img src="..."> with possible HTML tailings
    img_tag_matches = re.findall(
        r'<img[^>]+src=["\'](https?://[^"\'>]+)',  # stop at first quote or >
        text,
        re.IGNORECASE
    )
    
    # Match raw image URLs in plaintext
    raw_url_matches = re.findall(
        r'(https?://\S+\.(?:png|jpg|jpeg|gif|bmp)(?:\?\S*)?)',  # Allow query strings
        text,
        re.IGNORECASE
    )

    # Combine, decode, and clean URLs
    all_urls = img_tag_matches + raw_url_matches
    cleaned_urls = []

    for url in all_urls:
        # Unescape HTML entities like &amp;
        url = html.unescape(url.strip())

        # Remove any trailing HTML or unsafe characters
        url = re.sub(r'["\'<>\\]+$', '', url)

        cleaned_urls.append(url)

    return list(set(cleaned_urls))  # Remove duplicates


# Fetch image and convert to base64
def fetch_image_base64_from_url(image_url):
    try:
        print(f"[Downloading image from]: {image_url}")
        img_response = requests.get(image_url, auth=auth)  # Add auth here!
        if img_response.status_code == 200:
            encoded = base64.b64encode(img_response.content).decode('utf-8')
            print(f"[✓] Encoded image base64 length: {len(encoded)}")
            return encoded
        else:
            print(f"[✗] Failed to fetch image from {image_url} - Status Code: {img_response.status_code}")
            return None
    except Exception as e:
        print(f"[✗] Error fetching image: {e}")
        return None


# Function to fetch closed work items
def fetch_work_items():
    tickets = [] # temp list to collect tickets
    #response = requests.post(tfs_url, json=wiql_query, auth=auth, headers={'Content-Type': 'application/json'})
    response = request_with_retries(
        "post",
        tfs_url,
        json=wiql_query,
        auth=auth,
        headers={'Content-Type': 'application/json'}
    )

    if response.status_code == 200:
        work_items = response.json().get("workItems", [])
        if not work_items:
            print("No closed work items found in the last 14 days.")
        else:
            for item in work_items:
                work_item_id = item['id']
                print(f"Work Item ID: {work_item_id}")
                ticket = fetch_work_item_details(work_item_id)
                print(f"Fetched work ticket details!")
                if ticket:
                    print(f"Successfully added ticket to tickets")
                    tickets.append(ticket)
    else:
        print(f"Failed to fetch work items. Status Code: {response.status_code}")
        print(f"Error: {response.text}")
    
    return tickets

# URL to fetch details of a specific work item
base_url = "https://tfs-2018.casltd.com/tfs/Workpro/Workpro/_apis/wit/workitems/"

# Function to fetch work item details by ID
def fetch_work_item_details(work_item_id):
    url = f"{base_url}{work_item_id}?api-version=3.0"
    #response = requests.get(url, auth=auth)
    response = request_with_retries(
        "get",
        url,
        auth=auth
    )

    if response.status_code == 200:
        work_item_data = response.json()

        # Extract required fields dynamically
        title = work_item_data['fields'].get('System.Title', 'N/A')
        description = work_item_data['fields'].get('System.Description', 'N/A')
        internal_comments = work_item_data['fields'].get('Workpro.InternalComments', 'N/A')
        investigation_outcome = work_item_data['fields'].get('Workpro.InvestigationOutcome', 'N/A')
        root_cause = work_item_data['fields'].get('Workpro.RootCause', 'N/A')
        root_cause_reason = work_item_data['fields'].get('Workpro.RootCauseReason', 'N/A')
        how_fixed = work_item_data['fields'].get('Workpro.HowFixed', 'N/A')
        response_due_date = work_item_data['fields'].get('Workpro.ResponseDueDate', 'N/A')
        area = work_item_data['fields'].get('System.AreaPath', 'N/A')

        # Sanitise the description and internal comments fields
        '''Do we need to sanitise twice? We already sanitise in proces_and_upload.py'''
        sanitized_description = sanitize_text(description)
        sanitized_internal_comments = sanitize_text(internal_comments)

        # Try to extract and fetch image from description or internal comments
        image_urls = extract_all_image_urls(description) + extract_all_image_urls(internal_comments)
        print(f"[Image URLs] Found for {work_item_id}: {image_urls}")
        images_base64 = [fetch_image_base64_from_url(url) for url in image_urls if url]
        images_base64 = [img for img in images_base64 if img]  # Remove None values
        print(f"[Ticket Images] {work_item_id} -> {len(images_base64)} image(s) encoded.")

        # store in fields
        ticket = {
            "work_item_id" : work_item_id,
            "title" : title,
            "description": sanitized_description,
            "internal_comments": sanitized_internal_comments,
            "investigation_outcome": investigation_outcome,
            "root_cause": root_cause,
            "root_cause_reason": root_cause_reason,
            "how_fixed": how_fixed,
            "response_due_date": response_due_date,
            "area": area,
            "images": images_base64
        }

        return ticket
    else:
        print(f"Failed to fetch work item {work_item_id}, Status Code: {response.status_code}")
        return None #return none if failed

# Main function to orchestrate fetching work items and their details
def main():
    tickets = fetch_work_items()
    for t in tickets:
        print(t["work_item_id"], "-", t["title"])
    print(f"Successfully fetched tickets")
    return tickets


if __name__ == "__main__":
    all_tickets = main()

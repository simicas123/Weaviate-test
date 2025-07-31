import os
import requests
from requests.auth import HTTPBasicAuth
from datetime import datetime, timedelta
from sanitise_text import sanitize_text  # Import the sanitise function
import time
from requests.exceptions import RequestException, ConnectionError

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

        # Sanitise the description and internal comments fields
        sanitized_description = sanitize_text(description)
        sanitized_internal_comments = sanitize_text(internal_comments)

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

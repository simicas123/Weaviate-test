# script for fetching the items from TFS
import requests
from requests.auth import HTTPBasicAuth

ORG = "dev.azure.com/simi0872/" # azure login
PROJECT = "simi" # project name
PAT = "" # personal access token
 
auth = HTTPBasicAuth("", PAT) # getting the HTTP authentication using the PAT
 
def fetch_work_items():
    # Query to retrieve the work items from azure dev ops
    wiql = {
        "query": f"SELECT [System.Id] FROM WorkItems WHERE [System.TeamProject] = '{PROJECT}'"
    }
 
    wiql_url = f"https://{ORG}/{PROJECT}/_apis/wit/wiql?api-version=6.0"
    response = requests.post(wiql_url, json=wiql, auth=auth) # post the request and get the response
    response.raise_for_status()
    work_item_ids = [item["id"] for item in response.json()["workItems"]] # get work item ids from the response
 
    ids_str = ",".join(map(str, work_item_ids))
    items_url = f"https://{ORG}/_apis/wit/workitems?ids={ids_str}&api-version=6.0" #
    work_items = requests.get(items_url, auth=auth).json() #get request to fetch work item info
 
    for item in work_items["value"]:# loops through each returned work item
        print(item["fields"]["System.Title"]) # prints title of item
 
if __name__ == "__main__": 
    fetch_work_items() # runs fuunctionen script is executed
# script for fetching the items from TFS
import requests
from requests.auth import HTTPBasicAuth

ORG = "dev.azure.com/simi0872/" # azure login
PROJECT = "simi" # project name
PAT = "8N1poeKaOswP5VFGbrOd8FghScQnYwZCTs0ZbTOwtTwpDSTmQ4BxJQQJ99BGACAAAAApCpDPAAASAZDO2Ty5" # personal access token
 
auth = HTTPBasicAuth("", PAT) # getting the HTTP authentication using the PAT
 
def fetch_work_items():
    # Query to retrieve the work items from azure dev ops
    wiql = {
        "query": f"SELECT [System.Id] FROM WorkItems WHERE [System.TeamProject] = '{PROJECT}'"
    }
 
    wiql_url = f"https://{ORG}/{PROJECT}/_apis/wit/wiql?api-version=6.0"
    response = requests.post(wiql_url, json=wiql, auth=auth) # post the request and get the response
    response.raise_for_status()
    work_item_ids = [item["id"] for item in response.json()["workItems"]] # get work item ids from the response
 
    ids_str = ",".join(map(str, work_item_ids))
    items_url = f"https://{ORG}/_apis/wit/workitems?ids={ids_str}&api-version=6.0" 
    work_items = requests.get(items_url, auth=auth).json() #get request to fetch work item info
 
    for item in work_items["value"]:#loops through each returned work item
        print(item["fields"]["System.Title"]) #prints title of item
 
if __name__ == "__main__": 
    fetch_work_items() #runs functionen script is executed
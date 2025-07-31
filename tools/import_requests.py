import requests
 
proxy = {
    "http": "http://simi:JadeMoyo100%@172.16.0.171:8080",
    "https": "http://simi:JadeMoyo100%@172.16.0.171:8080"
}
 
try:
    response = requests.get("https://api.ipify.org", proxies=proxy, timeout=5)
    print("Proxy is working. IP returned:", response.text)
except requests.exceptions.ProxyError:
    print("Proxy error: Unable to connect.")
except requests.exceptions.ConnectTimeout:
    print("Connection timed out.")
except Exception as e:
    print("Other error:", e)

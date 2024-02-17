import json
import urllib.request

class SlackClient:
    def __init__(self, webhook_url: str):
        self.webhook_url = webhook_url

    def send_message(self, message: str) -> str:
        data = {
            "text": message
        }
        headers = {
            "Content-Type": "application/json"
        }

        # Create the request
        req = urllib.request.Request(
            self.webhook_url,
            data=json.dumps(data).encode("utf-8"),
            headers=headers
        )

        # Send the request
        try:
            with urllib.request.urlopen(req) as res:
                return res.read().decode("utf-8")
        except urllib.error.URLError as e:
            print("Failed to send message: ", e.reason)
            return ""

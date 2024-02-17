from api.slack import SlackClient
from dotenv import dotenv_values

# Load environment variables from the .env file
env_vars = dotenv_values(".env")
slack_client = SlackClient(env_vars.get("SLACK_WEBHOOK_URL"))
slack_client.send_message("test ririka.")

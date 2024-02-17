import sys
import os
sys.path.append(os.path.join(os.path.dirname(__file__), 'gen'))

from concurrent import futures
from dotenv import dotenv_values
import grpc
import gen.anime_radio_pb2
import gen.anime_radio_pb2_grpc
from api.slack import SlackClient

# Load environment variables from the .env file
env_vars = dotenv_values(".env")


class AnimeRadioService(gen.anime_radio_pb2_grpc.AnimeRadioServiceServicer):
    def SendAnimeRadioInfo(self, request_iterator, context):
        MAX_WORKERS = 10
        slack_client = SlackClient(env_vars.get("SLACK_WEBHOOK_URL"))

        # Create a thread pool
        with futures.ThreadPoolExecutor(MAX_WORKERS) as executor:
            # Submit the tasks
            future_to_url = {
                executor.submit(
                    slack_client.send_message,
                    f"title: {youtube_data.title}, URL: {youtube_data.url}"
                ): youtube_data
                for youtube_data in request_iterator
            }

            # Wait for the tasks to complete
            for future in futures.as_completed(future_to_url):
                youtube_data = future_to_url[future]
                try:
                    data = future.result()
                    print(f"Successfully sent message: {data}")
                except Exception as e:
                    print(f"Failed to send message: {e}")

        return gen.anime_radio_pb2.SlackResponse(
            result="success"
        )


def main():
    server = grpc.server(futures.ThreadPoolExecutor())
    gen.anime_radio_pb2_grpc.add_AnimeRadioServiceServicer_to_server(AnimeRadioService(), server)

    host = env_vars.get("GRPC_SERVER_HOST") or "localhost"
    port = env_vars.get("GRPC_SERVER_PORT") or 8080
    address = f"{host}:{port}"
    server.add_insecure_port(address)
    print(f"listen to {address}")
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    main()

import sys
import os
# Add the 'gen' directory to the Python path
sys.path.append(os.path.join(os.path.dirname(__file__), 'gen'))

from concurrent import futures
import grpc
import gen.anime_radio_pb2
import gen.anime_radio_pb2_grpc

class AnimeRadioService(gen.anime_radio_pb2_grpc.AnimeRadioServiceServicer):
    def SendAnimeRadioInfo(self, request_iterator, context):
        for youtube_data in request_iterator:
            print(f"title: {youtube_data.title}, URL: {youtube_data.url}")
        return gen.anime_radio_pb2.SlackResponse(
            result="success"
        )

def main():
    server = grpc.server(futures.ThreadPoolExecutor())
    gen.anime_radio_pb2_grpc.add_AnimeRadioServiceServicer_to_server(AnimeRadioService(), server)

    address = 'localhost:8080'
    server.add_insecure_port(address)
    print(f"listen to {address}")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    main()

package main

import (
	"context"
	"log"

	pb "github.com/miyuki-starmiya/anime-radio-grpc/gen"
	"google.golang.org/grpc"
)

// client stream
func SendAnimeRadioInfo(client pb.AnimeRadioServiceClient, dataItems []pb.YouTubeInfo) {
	stream, err := client.SendAnimeRadioInfo(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
	}

	for _, data := range dataItems {
		if err := stream.Send(&data); err != nil {
			log.Printf("Failed to send data: %v", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("Response: %s", response.GetResult())
}

func main() {
	// create connection to gRPC server
	conn, err := grpc.Dial(
		"localhost:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Printf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAnimeRadioServiceClient(conn)

	// mock data
	dataItems := []pb.YouTubeInfo{
		{Title: "title1", Url: "http://example.com/1"},
		{Title: "title2", Url: "http://example.com/2"},
	}
	SendAnimeRadioInfo(client, dataItems)
}

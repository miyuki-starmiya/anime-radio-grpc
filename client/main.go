package main

import (
	"context"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"

	"github.com/miyuki-starmiya/anime-radio-grpc/api"
	pb "github.com/miyuki-starmiya/anime-radio-grpc/gen"
	"github.com/miyuki-starmiya/anime-radio-grpc/variable"
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
	yc := api.NewYouTubeClient()
	dataItems, err := yc.SearchByKeyword(variable.KeywordList[0])
	log.Printf("dataItems: %v", dataItems)

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
	SendAnimeRadioInfo(client, dataItems)
}

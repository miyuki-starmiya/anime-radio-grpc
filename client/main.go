package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/miyuki-starmiya/anime-radio-grpc/api"
	"github.com/miyuki-starmiya/anime-radio-grpc/config"
	pb "github.com/miyuki-starmiya/anime-radio-grpc/gen"
	"github.com/miyuki-starmiya/anime-radio-grpc/variable"
)

const (
	retrySecond    = 30
	maxRetryMinute = 10
)

// Client stream
func SendAnimeRadioInfo(client pb.AnimeRadioServiceClient, dataItems []pb.YouTubeInfo, ctx context.Context) {
	stream, err := client.SendAnimeRadioInfo(ctx)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	for _, data := range dataItems {
		if err := stream.Send(&data); err != nil {
			log.Printf("Failed to send data: %v", err)
			return
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	log.Printf("Response: %s", response.GetResult())
}

func connectWithRetry(serverAddress string) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(maxRetryMinute)*time.Minute)
	defer cancel()

	for {
		conn, err = grpc.DialContext(ctx, serverAddress, grpc.WithInsecure())
		if err != nil {
			if status.Code(err) == codes.DeadlineExceeded {
				// Deadline exceeded, stop retrying
				log.Printf("Deadline %d minutes exceeded: %v", maxRetryMinute, err)
				return nil, err
			}
			// Retry
			log.Printf("Retry due to: %v", err)
			time.Sleep(time.Duration(retrySecond) * time.Second)
			continue
		}
		// Success
		break
	}
	return conn, nil
}

func main() {
	yc := api.NewYouTubeClient()

	var wg sync.WaitGroup
	keywords := append(variable.AnimeRadios, variable.VoiceActressRadios...)
	resChan := make(chan pb.YouTubeInfo, len(keywords))

	// Concurrently search YouTube data
	for _, keyword := range keywords {
		wg.Add(1)
		go func(kw string) {
			defer wg.Done()
			dataItems, err := yc.SearchByKeyword(kw)
			if err != nil {
				log.Printf("Error: %v", err)
				return
			}
			if len(dataItems) > 0 {
				resChan <- dataItems[0]
			}
		}(keyword)
	}

	// Wait for all goroutines until done
	go func() {
		wg.Wait()
		close(resChan)
	}()

	var youTubeInfos []pb.YouTubeInfo
	for data := range resChan {
		youTubeInfos = append(youTubeInfos, data)
	}
	log.Printf("youTubeInfos: %v", youTubeInfos)

	// Create connection to gRPC server
	var host string
	if os.getenv("ENV") == "production" {
		host = config.GRPC_SERVER_HOST
	} else {
		host = "localhost"
	}
	port := config.GRPC_SERVER_PORT
	serverAddress := fmt.Sprintf("%s:%s", host, port)
	conn, err := connectWithRetry(serverAddress)
	if err != nil {
		log.Printf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// After 10 minutes, cancel the client request
	client := pb.NewAnimeRadioServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(maxRetryMinute)*time.Minute)
	defer cancel()

	SendAnimeRadioInfo(client, youTubeInfos, ctx)
}

package api

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	pb "github.com/miyuki-starmiya/anime-radio-grpc/gen"
)

type YouTubeClient struct{}

func NewYouTubeClient() *YouTubeClient {
	return &YouTubeClient{}
}

func (yc *YouTubeClient) SearchByKeyword(keyword string) ([]pb.YouTubeInfo, error) {
	// create YouTube API client
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	service, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	// search by keyword
	call := service.Search.List([]string{"id", "snippet"}).Q(keyword).MaxResults(10)
	response, err := call.Do()
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	// create response
	var result []pb.YouTubeInfo
	for _, item := range response.Items {
		result = append(result, pb.YouTubeInfo{
			Title: item.Snippet.Title,
			Url:   "https://www.youtube.com/watch?v=" + item.Id.VideoId,
		})
		log.Printf("Title: %s, URL: %s", item.Snippet.Title, "https://www.youtube.com/watch?v="+item.Id.VideoId)
	}

	return result, nil
}

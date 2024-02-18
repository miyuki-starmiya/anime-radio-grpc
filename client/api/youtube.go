package api

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"

	"github.com/miyuki-starmiya/anime-radio-grpc/config"
	pb "github.com/miyuki-starmiya/anime-radio-grpc/gen"
)

type YouTubeClient struct {
	Service *youtube.Service
}

func NewYouTubeClient() *YouTubeClient {
	// init youtube client
	apiKey := config.YOUTUBE_API_KEY
	service, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Printf("Error: %v", err)
		return nil
	}

	return &YouTubeClient{
		Service: service,
	}
}

func (yc *YouTubeClient) SearchByKeyword(keyword string) ([]pb.YouTubeInfo, error) {
	// search by keyword
	// yesterdayStr := time.Now().Add(time.Duration(-24) * time.Hour).Format("2006-01-02T15:04:05Z")
	call := yc.Service.Search.List([]string{"id", "snippet"})
	// call = call.Q(keyword).MaxResults(1).PublishedAfter(yesterdayStr)
	call = call.Q(keyword).MaxResults(1)
	response, err := call.Do()
	if err != nil {
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

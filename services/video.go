package services

import (
	"context"
	"server/configs"
	"server/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type VideoService struct {
	videoCollection *mongo.Collection
}

func NewVideoService() *VideoService {
	videoService := new(VideoService)
	videoService.videoCollection = configs.GetCollection(configs.DB, "video")

	return videoService
}

func (videoService *VideoService) CreateVideo(ctx context.Context, video models.Video) *models.Video {
	_, err := videoService.videoCollection.InsertOne(ctx, video)
	if err != nil {
		return nil
	}

	return &video
}

func (videoService *VideoService) GetVideos(ctx context.Context, number int, page int) []models.Video {
	var videos []models.Video
	cursor, err := videoService.videoCollection.Find(ctx, &fiber.Map{})
	if err != nil {
		return nil
	}
	for cursor.Next(ctx) {
		var result models.Video
		if err := cursor.Decode(&result); err != nil {
			return nil
		}
		videos = append(videos, result)
	}
	return videos

}

package services

import (
	"context"
	"server/configs"
	"server/models"

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

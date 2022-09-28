package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/dto"
	"server/middleware"
	"server/models"
	"server/services"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VideoController struct {
	userService    services.UserService
	accountService services.AccountService
	videoService   services.VideoService
}

func NewVideoController() *VideoController {
	videoController := new(VideoController)
	videoController.accountService = *services.NewAccountService()
	videoController.userService = *services.NewUserService()
	videoController.videoService = *services.NewVideoService()

	return videoController
}

func (videoController *VideoController) VideoApi(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	headerAuthorization := c.GetReqHeaders()["Authorization"]
	token := strings.Split(headerAuthorization, " ")[1]
	if token == "" {
		return c.Status(http.StatusAccepted).JSON(dto.ResponseCreateVideoDTO{Success: false})
	}
	accountId, _ := primitive.ObjectIDFromHex(middleware.GetAccountId(token))

	account := videoController.accountService.GetAccountById(ctx, accountId)
	if account == nil {
		return c.Status(http.StatusAccepted).JSON(dto.ResponseCreateVideoDTO{Success: false})
	}

	var createVideoDTO dto.CreateVideoDTO

	if err := c.BodyParser(&createVideoDTO); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseCreateVideoDTO{Success: false})
	}

	id := primitive.NewObjectID()

	fileHeader, err := c.FormFile("video")
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseCreateVideoDTO{Success: false})
	}

	file, _ := fileHeader.Open()
	data, _ := ioutil.ReadAll(file)

	postFix := ".mp4"

	if strings.Split(http.DetectContentType(data), "/")[0] != "video" {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseCreateVideoDTO{Success: false, Message: "Not Allow Anything else Video File"})
	}

	err = ioutil.WriteFile("./storage/video/"+id.Hex()+postFix, data, 0644)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseCreateVideoDTO{Success: false})
	}

	video := models.Video{
		Id:          id,
		Size:        strconv.Itoa(int(fileHeader.Size)) + " bytes",
		UserId:      account.UserId,
		CreatedAt:   time.Now(),
		Description: createVideoDTO.Description,
	}
	addedVideo := videoController.videoService.CreateVideo(ctx, video)

	if addedVideo == nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseCreateVideoDTO{Success: false})
	}

	return c.Status(http.StatusAccepted).JSON(dto.ResponseCreateVideoDTO{Success: true, Video: *addedVideo})
}

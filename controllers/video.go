package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"server/dto"
	"server/models"
	"server/services"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VideoController struct {
	userService    *services.UserService
	accountService *services.AccountService
	videoService   *services.VideoService
	channelService *services.ChannelService
}

func NewVideoController() *VideoController {
	videoController := new(VideoController)
	videoController.accountService = services.NewAccountService()
	videoController.userService = services.NewUserService()
	videoController.videoService = services.NewVideoService()
	videoController.channelService = services.NewChannelService()

	return videoController
}

func (videoController *VideoController) VideoApi(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	accountId, _ := primitive.ObjectIDFromHex(c.GetReqHeaders()["Id"])

	account := videoController.accountService.GetAccountById(ctx, accountId)

	if account == nil {
		return c.Status(400).JSON(dto.ResponseCreateVideoDTO{Success: false, Message: "Undentified Token"})
	}

	channel := videoController.channelService.GetChannelByUserId(ctx, account.UserId)

	if channel == nil {
		return c.Status(400).JSON(dto.ResponseCreateVideoDTO{Success: false, Message: "Invalid Channel"})
	}

	var createVideoDTO dto.CreateVideoDTO

	if err := c.BodyParser(&createVideoDTO); err != nil {
		return c.Status(500).JSON(dto.ResponseCreateVideoDTO{Success: false, Message: "Internal Error Server"})
	}

	id := primitive.NewObjectID()

	fileHeader, err := c.FormFile("video")
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(dto.ResponseCreateVideoDTO{Success: false, Message: "Internal Error Server"})
	}

	file, _ := fileHeader.Open()
	data, _ := ioutil.ReadAll(file)

	postFix := ".mp4"

	if strings.Split(http.DetectContentType(data), "/")[0] != "video" {
		return c.Status(400).JSON(dto.ResponseCreateVideoDTO{Success: false, Message: "Not Allow Anything else Video File"})
	}
	if err := os.Mkdir("./storage/video/"+id.Hex(), 0755); err != nil {
		return c.Status(500).JSON(dto.ResponseCreateVideoDTO{Success: false, Message: "Internal Error Server"})
	}
	err = ioutil.WriteFile("./storage/video/"+id.Hex()+"/video"+postFix, data, 0644)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(dto.ResponseCreateVideoDTO{Success: false, Message: "Internal Error Server"})
	}

	channel = videoController.channelService.GetChannelByUserId(ctx, account.UserId)

	if channel == nil {
		return c.Status(400).JSON(dto.ResponseCreateVideoDTO{Success: false, Message: "Not Channel Error"})
	}

	video := models.Video{
		Id:          id,
		Name:        createVideoDTO.Name,
		Size:        strconv.Itoa(int(fileHeader.Size)) + " bytes",
		ChannelId:   channel.Id,
		CreatedAt:   time.Now(),
		Description: createVideoDTO.Description,
		View:        0,
	}
	addedVideo := videoController.videoService.CreateVideo(ctx, video)

	if addedVideo == nil {
		return c.Status(500).JSON(dto.ResponseCreateVideoDTO{Success: false, Message: "Internal Error Server"})
	}

	return c.Status(200).JSON(dto.ResponseCreateVideoDTO{Success: true, Video: *addedVideo})
}

func (videoController *VideoController) GetVideos(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	number := c.Query("number")
	page := c.Query("page")

	defer cancel()
	num := 10
	if numQuery, err := strconv.Atoi(number); err == nil {
		num = numQuery
	}

	pg := 0

	if pageQuery, err := strconv.Atoi(page); err == nil {
		pg = pageQuery
	}

	videos := videoController.videoService.GetVideos(ctx, num, pg)

	var users []dto.InfoUserUploadedVideo

	for _, video := range videos {
		channel := videoController.channelService.GetChannelById(ctx, video.ChannelId)
		user := videoController.userService.GetUserById(ctx, channel.UserId)
		users = append(users, dto.InfoUserUploadedVideo{Id: user.Id.Hex(), Name: user.FirstName + " " + user.LastName})
	}

	return c.Status(200).JSON(&fiber.Map{"videos": videos, "users": users})
}

func (videoController *VideoController) GetMedia(c *fiber.Ctx) error {
	id := c.Query("id")

	return c.Status(200).SendFile("./storage/video/" + id + "/video.mp4")
}

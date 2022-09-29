package controllers

import (
	"server/services"
)

type ChannelController struct {
	userService    *services.UserService
	accountService *services.AccountService
	channelService *services.ChannelService
}

func NewChannelController() *ChannelController {
	channelController := new(ChannelController)
	channelController.accountService = services.NewAccountService()
	channelController.userService = services.NewUserService()
	channelController.channelService = services.NewChannelService()
	return channelController
}

/* func (channelController *ChannelController) CreateChannel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()


}
*/

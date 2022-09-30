package dto

type CreateSubscriberDTO struct {
	ChannelId string `json:"channelId"`
}

type ResponseCreateSubscriberDTO struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

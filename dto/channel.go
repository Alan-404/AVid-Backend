package dto

import "server/models"

type ResponseCreateChannel struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Channel *models.Channel `json:"channel"`
}

type ResponseGetChannelById struct {
	Success bool            `json:"success"`
	Channel *models.Channel `json:"channel"`
}

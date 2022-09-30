package dto

import (
	"os"
	"server/models"
)

type CreateVideoDTO struct {
	Video       os.File `form:"video"`
	Name        string  `form:"name"`
	Description string  `form:"description"`
}

type ResponseCreateVideoDTO struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Video   models.Video `json:"video"`
}

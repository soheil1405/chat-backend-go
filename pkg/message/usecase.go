package message

import (
	"chat-app/models"

	"github.com/labstack/echo/v4"
)

type IMessageUC interface {
	SendMessage(ctx echo.Context, message *models.Message) (*models.Message, error)
}

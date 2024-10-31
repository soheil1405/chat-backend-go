package usecase

import (
	"chat-app/models"
	"chat-app/pkg/message"

	"github.com/labstack/echo/v4"
)

type MessageUC struct {
	MessageRepo message.IMessageRepo
}

func NewMessageUC(repo message.IMessageRepo) message.IMessageUC {
	return &MessageUC{
		MessageRepo: repo,
	}

}

func (m *MessageUC) SendMessage(ctx echo.Context, message *models.Message) (*models.Message, error) {
	return m.MessageRepo.SendMessage(message)
}

package message

import "chat-app/models"

type IMessageRepo interface {
	SendMessage(message *models.Message) (*models.Message, error)
}

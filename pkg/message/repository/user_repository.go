package repository

import (
	"chat-app/models"

	"gorm.io/gorm"
)

type MessageRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *MessageRepo {
	return &MessageRepo{
		db: db,
	}
}

func (m *MessageRepo) SendMessage(message *models.Message) (*models.Message, error) {
	if err := m.db.Create(&message).Error; err != nil {
		return nil, err
	}

	return message, nil
}

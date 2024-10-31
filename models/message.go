package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Message   string `json:"message,omitempty"`
	SenderID  uint   `json:"sender_id,omitempty"`
	ReciverID uint   `json:"reciver_id,omitempty"`
}

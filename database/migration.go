package database

import (
	"chat-app/models"
	"fmt"

	"gorm.io/gorm"
)

func DoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Message{}); err != nil {
		fmt.Println("err in migrating messages")
	}
	fmt.Println("migration successed")
}

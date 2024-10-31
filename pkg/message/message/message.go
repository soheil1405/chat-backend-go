package message

import (
	"chat-app/database"
	"chat-app/pkg/message/delivery/http"
	"chat-app/pkg/message/repository"
	"chat-app/pkg/message/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, e echo.Group) {
	repo := repository.New(db)
	usecase := usecase.NewMessageUC(repo)
	http.New(usecase, e)
	database.DoMigrate(db)

}

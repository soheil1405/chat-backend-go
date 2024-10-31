package main

import (
	"chat-app/database"
	"chat-app/pkg/message/message"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {
	var (
		db *gorm.DB
		e  = echo.New()
	)

	db, _ = database.Connection()
	fmt.Println(db)
	routes := e.Group("/api/v1")
	message.Init(db, *routes)
	e.Logger.Fatal(e.Start(":1323"))

}

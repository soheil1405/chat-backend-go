package http

import (
	"chat-app/models"
	"chat-app/pkg/message"
	"errors"

	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	uc message.IMessageUC
}

func New(messageUC message.IMessageUC, router echo.Group) {
	handler := MessageHandler{
		uc: messageUC,
	}

	messageRouter := router.Group("/messages")
	messageRouter.POST("/send", handler.sendMessage)

}

func (h MessageHandler) sendMessage(ctx echo.Context) (err error) {
	var (
		message *models.Message
	)
	if err = ctx.Bind(message); err != nil {
		return ctx.JSON(400, errors.New("bad request"))
	}

	if message, err = h.uc.SendMessage(ctx, message); err != nil {
		return ctx.JSON(400, err)
	}
	return ctx.JSON(200, message)
}

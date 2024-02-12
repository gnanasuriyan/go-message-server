package services

import (
	"github.com/gnanasuriyan/go-message-server/app/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type IMessageService interface {
	PostMessage(ctx *fiber.Ctx) error
}

type MessageService struct {
	MessageRepository repositories.IMessageRepository
}

var NewMessageService = wire.NewSet(wire.Struct(new(MessageService), "*"), wire.Bind(new(IMessageService), new(*MessageService)))

func (s *MessageService) PostMessage(ctx *fiber.Ctx) error {
	return nil
}

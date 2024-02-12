package services

import (
	"strconv"

	"github.com/gnanasuriyan/go-message-server/app/consts"

	"github.com/gnanasuriyan/go-message-server/app/models"
	"github.com/gnanasuriyan/go-message-server/app/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type IMessageService interface {
	PostMessage(ctx *fiber.Ctx) error
	ListMessages(ctx *fiber.Ctx) error
}

type MessageService struct {
	MessageRepository repositories.IMessageRepository
	UserRepository    repositories.IUserRepository
}

var NewMessageService = wire.NewSet(wire.Struct(new(MessageService), "*"), wire.Bind(new(IMessageService), new(*MessageService)))

func (s *MessageService) PostMessage(ctx *fiber.Ctx) error {
	messageCreateDto := new(models.MessageCreateDto)
	if err := ctx.BodyParser(messageCreateDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, consts.InvalidRequestBody)
	}
	userIdStr := ctx.GetReqHeaders()["X-Auth-User-Id"]
	if userIdStr == nil || len(userIdStr) == 0 {
		return fiber.NewError(fiber.StatusUnauthorized, consts.Unauthorized)
	}
	userId, err := strconv.Atoi(userIdStr[0])
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, consts.Unauthorized)
	}
	message, err := s.MessageRepository.Insert(ctx, uint(userId), *messageCreateDto)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, consts.SomethingWentWrong)
	}
	return ctx.JSON(message)
}

func (s *MessageService) ListMessages(ctx *fiber.Ctx) error {
	userIdStr := ctx.GetReqHeaders()["X-Auth-User-Id"]
	if userIdStr == nil || len(userIdStr) == 0 {
		return fiber.NewError(fiber.StatusUnauthorized, consts.Unauthorized)
	}
	currentUserId, err := strconv.Atoi(userIdStr[0])
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, consts.Unauthorized)
	}
	paginationDto := new(models.PaginationDto)
	if err := ctx.QueryParser(paginationDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, consts.InvalidQueryParameters)
	}
	messages, err := s.MessageRepository.FindAll(ctx, *paginationDto)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, consts.SomethingWentWrong)
	}
	var messageResponse []models.MessageResponseDto
	// TODO: this can be optimized by using a join query or cache the user details
	for _, message := range messages {
		user, err := s.UserRepository.UserById(ctx, int(message.FkUser))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, consts.SomethingWentWrong)
		}
		messageResponse = append(messageResponse, models.MessageResponseDto{
			ID:                    message.ID,
			Content:               message.Content,
			CreatedAt:             message.CreatedAt,
			PostedBy:              user.Username,
			IsPostedByCurrentUser: currentUserId == int(message.FkUser),
		})
	}
	return ctx.JSON(messageResponse)
}

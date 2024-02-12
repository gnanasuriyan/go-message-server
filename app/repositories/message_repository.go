package repositories

import (
	"github.com/gofiber/fiber/v2"

	"github.com/google/wire"

	"github.com/gnanasuriyan/go-message-server/internal/db"

	"github.com/gnanasuriyan/go-message-server/app/models"
)

type IMessageRepository interface {
	Insert(ctx *fiber.Ctx, userId uint, dto models.MessageCreateDto) (*models.Message, error)
	FindAll(ctx *fiber.Ctx, pagination models.PaginationDto) ([]models.Message, error)
}

type MessageRepository struct {
	Db db.IAppDB
}

var NewMessageRepository = wire.NewSet(wire.Struct(new(MessageRepository), "*"), wire.Bind(new(IMessageRepository), new(*MessageRepository)))

func (r *MessageRepository) FindAll(ctx *fiber.Ctx, pagination models.PaginationDto) ([]models.Message, error) {
	var messages []models.Message
	tx := r.Db.Where("`active` = ?", 1).Limit(pagination.Limit).Offset(pagination.Limit * (pagination.Page - 1)).Find(&messages)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return messages, nil
}

func (r *MessageRepository) Insert(ctx *fiber.Ctx, userId uint, dto models.MessageCreateDto) (*models.Message, error) {
	message := models.Message{
		FkUser:  userId,
		Content: dto.Content,
		Active:  true,
	}
	tx := r.Db.Create(&message)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &message, nil
}

package repositories

import (
	"context"

	"github.com/google/wire"

	"github.com/gnanasuriyan/go-message-server/internal/db"

	"github.com/gnanasuriyan/go-message-server/app/models"
)

type IMessageRepository interface {
	Insert(ctx context.Context, dto models.MessageCreate) (*models.Message, error)
	FindAll(ctx context.Context, pagination models.Pagination) ([]models.Message, error)
}

type MessageRepository struct {
	Db db.IAppDB
}

var NewMessageRepository = wire.NewSet(wire.Struct(new(MessageRepository), "*"), wire.Bind(new(IMessageRepository), new(*MessageRepository)))

func (r *MessageRepository) FindAll(ctx context.Context, pagination models.Pagination) ([]models.Message, error) {
	var messages []models.Message
	tx := r.Db.Where("`active` = ?", 1).Limit(pagination.Limit).Offset(pagination.Limit * (pagination.Page - 1)).Find(&messages)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return messages, nil
}

func (r *MessageRepository) Insert(ctx context.Context, dto models.MessageCreate) (*models.Message, error) {
	message := models.Message{
		FkUser:  dto.FkUser,
		Content: dto.Content,
		Active:  true,
	}
	tx := r.Db.Create(&message)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &message, nil
}

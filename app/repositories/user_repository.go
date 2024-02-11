package repositories

import (
	"context"

	"github.com/gnanasuriyan/go-message-server/internal/db"

	"github.com/gnanasuriyan/go-message-server/app/models"
	"github.com/google/wire"
)

type IUserRepository interface {
	Insert(ctx context.Context, user models.User) (*models.User, error)
	UserById(ctx context.Context, id int) (*models.User, error)
	UserByUserName(ctx context.Context, username string) (*models.User, error)
}

type UserRepository struct {
	Db db.IAppDB
}

var NewUserRepository = wire.NewSet(wire.Struct(new(UserRepository), "*"), wire.Bind(new(IUserRepository), new(*UserRepository)))

func (r *UserRepository) Insert(ctx context.Context, user models.User) (*models.User, error) {
	tx := r.Db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (r *UserRepository) UserById(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	tx := r.Db.Where("`id` = ? AND `active` = ?", id, 1).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (r *UserRepository) UserByUserName(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	tx := r.Db.Where("`username` = ? AND `active` = ?", username, 1).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

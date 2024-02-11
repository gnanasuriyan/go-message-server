package repositories

import (
	"context"

	"github.com/gnanasuriyan/go-message-server/app/models"
	"github.com/google/wire"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.UserCreate) (*models.User, error)
	UserById(ctx context.Context, id int) (*models.User, error)
	UserByUserName(ctx context.Context, email string) (*models.User, error)
}

type UserRepository struct{}

var NewUserRepository = wire.NewSet(wire.Struct(new(UserRepository), "*"), wire.Bind(new(IUserRepository), new(*UserRepository)))

func (r *UserRepository) UserById(ctx context.Context, id int) (*models.User, error) {
	return &models.User{}, nil
}

func (r *UserRepository) UserByUserName(ctx context.Context, email string) (*models.User, error) {
	return &models.User{}, nil
}

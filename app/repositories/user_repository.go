package repositories

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gnanasuriyan/go-message-server/internal/db"

	"github.com/gnanasuriyan/go-message-server/app/models"
	"github.com/google/wire"
)

type IUserRepository interface {
	Insert(ctx *fiber.Ctx, user models.User) (*models.User, error)
	UserById(ctx *fiber.Ctx, id int) (*models.User, error)
	UserByUserName(ctx *fiber.Ctx, username string) (*models.User, error)
}

type UserRepository struct {
	Db db.IAppDB
}

var NewUserRepository = wire.NewSet(wire.Struct(new(UserRepository), "*"), wire.Bind(new(IUserRepository), new(*UserRepository)))

func (r *UserRepository) Insert(ctx *fiber.Ctx, user models.User) (*models.User, error) {
	tx := r.Db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (r *UserRepository) UserById(ctx *fiber.Ctx, id int) (*models.User, error) {
	var user models.User
	tx := r.Db.Where("`id` = ? AND `active` = ?", id, 1).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (r *UserRepository) UserByUserName(ctx *fiber.Ctx, username string) (*models.User, error) {
	var user models.User
	tx := r.Db.Where("`username` = ? AND `active` = ?", username, 1).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

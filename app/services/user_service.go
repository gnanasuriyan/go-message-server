package services

import (
	"github.com/gnanasuriyan/go-message-server/app/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type IUserService interface {
	CreateNewUser(ctx *fiber.Ctx) error
	AuthenticateUser(ctx *fiber.Ctx) error
}

type UserService struct {
	UserRepository repositories.IUserRepository
}

var NewUserService = wire.NewSet(wire.Struct(new(UserService), "*"), wire.Bind(new(IUserService), new(*UserService)))

func (s *UserService) CreateNewUser(ctx *fiber.Ctx) error {
	return nil
}

func (s *UserService) AuthenticateUser(ctx *fiber.Ctx) error {
	return nil
}

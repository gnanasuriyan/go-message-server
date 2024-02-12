package services

import (
	"strings"

	"github.com/gnanasuriyan/go-message-server/app/models"
	"github.com/gnanasuriyan/go-message-server/app/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type IUserService interface {
	Signup(ctx *fiber.Ctx) error
	AuthenticateUser(ctx *fiber.Ctx) error
}

type UserService struct {
	UserRepository repositories.IUserRepository
}

var NewUserService = wire.NewSet(wire.Struct(new(UserService), "*"), wire.Bind(new(IUserService), new(*UserService)))

func (s *UserService) Signup(ctx *fiber.Ctx) error {
	userCreateDto := new(models.UserCreateDto)
	if err := ctx.BodyParser(userCreateDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	if strings.Trim(userCreateDto.Password, " ") != strings.Trim(userCreateDto.ConfirmPassword, " ") {
		return fiber.NewError(fiber.StatusBadRequest, "Password and Confirm Password are not same")
	}
	//TODO: password hash instead of raw password
	_, err := s.UserRepository.Insert(ctx, models.User{
		Username: userCreateDto.Username,
		Password: userCreateDto.Password,
		Active:   true,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong while creating user")
	}
	return ctx.JSON(models.SuccessResponseDto{
		Success: true,
		Message: "User signed up successfully",
	})
}

func (s *UserService) AuthenticateUser(ctx *fiber.Ctx) error {
	authenticateUserDto := new(models.AuthenticateUserDto)
	if err := ctx.BodyParser(authenticateUserDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	user, err := s.UserRepository.UserByUserName(ctx, authenticateUserDto.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong while authenticating user")
	}
	if user == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
	}
	if user.Password != authenticateUserDto.Password {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
	}
	// TODO: return JWT token
	return ctx.JSON(user)
}

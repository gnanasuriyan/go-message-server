package services

import (
	"strings"

	"github.com/gnanasuriyan/go-message-server/app/consts"

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
		return fiber.NewError(fiber.StatusBadRequest, consts.InvalidRequestBody)
	}
	if strings.Trim(userCreateDto.Password, " ") != strings.Trim(userCreateDto.ConfirmPassword, " ") {
		return fiber.NewError(fiber.StatusBadRequest, consts.PasswordMismatch)
	}
	//TODO: save password hash instead of raw password
	_, err := s.UserRepository.Insert(ctx, models.User{
		Username: userCreateDto.Username,
		Password: userCreateDto.Password,
		Active:   true,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, consts.SomethingWentWrong)
	}
	return ctx.JSON(models.SuccessResponseDto{
		Success: true,
		Message: "User signed up successfully",
	})
}

func (s *UserService) AuthenticateUser(ctx *fiber.Ctx) error {
	loginRequestDto := new(models.LoginRequestDto)
	if err := ctx.BodyParser(loginRequestDto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, consts.InvalidRequestBody)
	}
	user, err := s.UserRepository.UserByUserName(ctx, loginRequestDto.Username)
	if err != nil && err.Error() == consts.RecordNotFound {
		return fiber.NewError(fiber.StatusUnauthorized, consts.InvalidUsernameOrPassword)
	} else if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, consts.SomethingWentWrong)
	}
	if user.Password != loginRequestDto.Password {
		return fiber.NewError(fiber.StatusUnauthorized, consts.InvalidUsernameOrPassword)
	}
	// TODO: return JWT token
	return ctx.JSON(models.LoginUserDto{
		Username: user.Username,
		ID:       user.ID,
	})
}

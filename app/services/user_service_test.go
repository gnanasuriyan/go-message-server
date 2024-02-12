package services_test

import (
	"errors"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/gnanasuriyan/go-message-server/app/services"
	"github.com/stretchr/testify/assert"

	"github.com/gnanasuriyan/go-message-server/app/repositories"
	"github.com/gnanasuriyan/go-message-server/test"

	"github.com/gofiber/fiber/v2"
)

func TestUserService_Signup_Wrong_Credentials(t *testing.T) {
	app := fiber.New()
	dbMock := test.NewMockDB()
	userRepository := repositories.UserRepository{Db: dbMock.DB}
	userService := services.UserService{UserRepository: &userRepository}
	app.Post("/", userService.Signup)

	reqBody := `{
		"username": "test_user_1",
		"password": "123",
		"confirm_password": "1234"
	}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)
	a := assert.New(t)
	a.NotNil(res)
	a.Equal(400, res.StatusCode)
}

func TestUserService_Signup_Database_Error(t *testing.T) {
	app := fiber.New()
	dbMock := test.NewMockDB()
	userRepository := repositories.UserRepository{Db: dbMock.DB}
	mock := dbMock.Sqlmock
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `users` (.+)$").WillReturnError(errors.New("error"))
	mock.ExpectRollback()
	userService := services.UserService{UserRepository: &userRepository}
	app.Post("/", userService.Signup)

	reqBody := `{
		"username": "test_user_1",
		"password": "123",
		"confirm_password": "123"
	}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)
	a := assert.New(t)
	a.NotNil(res)
	a.Equal(500, res.StatusCode)
}

func TestUserService_Signup_Happy_Flow(t *testing.T) {
	app := fiber.New()
	dbMock := test.NewMockDB()
	userRepository := repositories.UserRepository{Db: dbMock.DB}
	mock := dbMock.Sqlmock
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `users` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	userService := services.UserService{UserRepository: &userRepository}
	app.Post("/", userService.Signup)

	reqBody := `{
		"username": "test_user_1",
		"password": "123",
		"confirm_password": "123"
	}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)
	a := assert.New(t)
	a.NotNil(res)
	a.Equal(200, res.StatusCode)
}

func TestUserService_AuthenticateUser(t *testing.T) {
	app := fiber.New()
	dbMock := test.NewMockDB()
	userRepository := repositories.UserRepository{Db: dbMock.DB}
	mock := dbMock.Sqlmock
	rows := sqlmock.NewRows([]string{"id", "username", "password", "active", "shadow_active", "created_at", "updated_at"}).AddRow(1, "test_user_1", "123", true, true, time.Now(), time.Now())
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `username` = ? AND `active` = ? ORDER BY `users`.`id` LIMIT ?")).WillReturnRows(rows)
	userService := services.UserService{UserRepository: &userRepository}
	app.Post("/", userService.AuthenticateUser)
	reqBody := `{
		"username": "test_user_1",
		"password": "123"
	}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)
	a := assert.New(t)
	a.NotNil(res)
	a.Equal(200, res.StatusCode)
}

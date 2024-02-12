package services_test

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gnanasuriyan/go-message-server/app/repositories"

	"github.com/gnanasuriyan/go-message-server/app/services"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gnanasuriyan/go-message-server/test"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestMessageService_PostMessage_No_Error(t *testing.T) {
	dbMock := test.NewMockDB()
	messageRepository := repositories.MessageRepository{Db: dbMock.DB}
	mock := dbMock.Sqlmock
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `messages` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	app := fiber.New()
	messageService := services.MessageService{MessageRepository: &messageRepository}
	app.Post("/", messageService.PostMessage)
	reqBody := `{
		"content": "Sample message 1",
		"fk_user": 1
	}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)
	a := assert.New(t)
	a.NotNil(res)
	a.Equal(200, res.StatusCode)
}

func TestMessageService_PostMessage_With_DB_Error(t *testing.T) {
	dbMock := test.NewMockDB()
	messageRepository := repositories.MessageRepository{Db: dbMock.DB}
	mock := dbMock.Sqlmock
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `messages` (.+)$").WillReturnError(errors.New("error"))
	mock.ExpectRollback()

	app := fiber.New()
	messageService := services.MessageService{MessageRepository: &messageRepository}
	app.Post("/", messageService.PostMessage)
	reqBody := `{
		"content": "Sample message 1",
		"fk_user": 1
	}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)
	a := assert.New(t)
	a.NotNil(res)
	a.Equal(500, res.StatusCode)
}

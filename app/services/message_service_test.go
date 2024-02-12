package services_test

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/valyala/fasthttp"

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
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	c.Set("user_id", "1")
	messageService := services.MessageService{MessageRepository: &messageRepository}
	app.Post("/", messageService.PostMessage)
	reqBody := `{
		"content": "Sample message 1"
	}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-User-Id", "1")
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
		"content": "Sample message 1"
	}`
	req := httptest.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-User-Id", "1")
	res, _ := app.Test(req)
	a := assert.New(t)
	a.NotNil(res)
	a.Equal(500, res.StatusCode)
}

func TestMessageService_PostMessage_ListMessages(t *testing.T) {
	dbMock := test.NewMockDB()
	mock := dbMock.Sqlmock
	rows := sqlmock.NewRows([]string{"id", "fk_user", "content", "active", "created_at", "updated_at"}).
		AddRow(1, 1, "Sample message 1", true, time.Now(), time.Now()).
		AddRow(2, 1, "Sample message 2", true, time.Now(), time.Now()).
		AddRow(3, 1, "Sample message 3", true, time.Now(), time.Now()).
		AddRow(4, 1, "Sample message 4", true, time.Now(), time.Now()).
		AddRow(5, 1, "Sample message 5", true, time.Now(), time.Now()).
		AddRow(6, 1, "Sample message 6", true, time.Now(), time.Now()).
		AddRow(7, 1, "Sample message 7", true, time.Now(), time.Now()).
		AddRow(8, 1, "Sample message 8", true, time.Now(), time.Now()).
		AddRow(9, 1, "Sample message 9", true, time.Now(), time.Now()).
		AddRow(10, 1, "Sample message 10", true, time.Now(), time.Now())

	messageRepository := repositories.MessageRepository{Db: dbMock.DB}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `messages` WHERE `active` = ? LIMIT ?")).WillReturnRows(rows)
	app := fiber.New()
	messageService := services.MessageService{MessageRepository: &messageRepository}
	app.Get("/", messageService.ListMessages)
	req := httptest.NewRequest("GET", fmt.Sprintf("/?page=%d&limit=%d", 1, 10), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-User-Id", "1")
	res, _ := app.Test(req)
	a := assert.New(t)
	a.NotNil(res)
	a.Equal(200, res.StatusCode)
}

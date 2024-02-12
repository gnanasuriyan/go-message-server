package repositories_test

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/gnanasuriyan/go-message-server/app/repositories"

	"github.com/gnanasuriyan/go-message-server/app/models"

	"github.com/gnanasuriyan/go-message-server/test"
)

func TestUserRepository_Insert_No_Error(t *testing.T) {
	dbMock := test.NewMockDB()
	userRepository := repositories.UserRepository{Db: dbMock.DB}
	mock := dbMock.Sqlmock
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `users` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	user, err := userRepository.Insert(ctx, models.User{
		Username: "test_user_1",
		Password: "123",
		Active:   true,
	})
	a := assert.New(t)
	a.NoError(err)
	a.NotNil(user)
}

func TestUserRepository_Insert_Return_Error(t *testing.T) {
	dbMock := test.NewMockDB()
	userRepository := repositories.UserRepository{Db: dbMock.DB}
	mock := dbMock.Sqlmock
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `users` (.+)$").WillReturnError(errors.New("error"))
	mock.ExpectRollback()

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	user, err := userRepository.Insert(ctx, models.User{
		Username: "test_user_1",
		Password: "123",
		Active:   true,
	})
	a := assert.New(t)
	a.Error(err)
	a.Nil(user)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_UserById(t *testing.T) {
	dbMock := test.NewMockDB()
	userRepository := repositories.UserRepository{Db: dbMock.DB}
	mock := dbMock.Sqlmock
	rows := sqlmock.NewRows([]string{"id", "username", "password", "active", "shadow_active", "created_at", "updated_at"}).AddRow(1, "test_user_1", "123", true, true, time.Now(), time.Now())
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `id` = ? AND `active` = ? ORDER BY `users`.`id` LIMIT ?")).WillReturnRows(rows)

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	user, err := userRepository.UserById(ctx, 1)
	a := assert.New(t)
	a.NoError(err)
	a.NotNil(user)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserRepository_UserByUserName(t *testing.T) {
	dbMock := test.NewMockDB()
	userRepository := repositories.UserRepository{Db: dbMock.DB}
	mock := dbMock.Sqlmock
	rows := sqlmock.NewRows([]string{"id", "username", "password", "active", "shadow_active", "created_at", "updated_at"}).AddRow(1, "test_user_1", "123", true, true, time.Now(), time.Now())
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `username` = ? AND `active` = ? ORDER BY `users`.`id` LIMIT ?")).WillReturnRows(rows)

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	user, err := userRepository.UserByUserName(ctx, "test_user_1")
	a := assert.New(t)
	a.NoError(err)
	a.NotNil(user)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

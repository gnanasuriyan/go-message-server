package repositories_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gnanasuriyan/go-message-server/app/models"
	"github.com/gnanasuriyan/go-message-server/app/repositories"
	"github.com/gnanasuriyan/go-message-server/test"
	"github.com/stretchr/testify/assert"
)

func TestMessageRepository_FindAll(t *testing.T) {
	ctx := context.Background()
	dbMock := test.NewMockDB()
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
	mock := dbMock.Sqlmock
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `messages` WHERE `active` = ? LIMIT ?")).WillReturnRows(rows)
	messages, err := messageRepository.FindAll(ctx, models.Pagination{
		Page:  1,
		Limit: 10,
	})
	assert := assert.New(t)
	assert.NoError(err)
	assert.NotNil(messages)
}

func TestMessageRepository_Insert_No_Error(t *testing.T) {
	ctx := context.Background()
	dbMock := test.NewMockDB()
	messageRepository := repositories.MessageRepository{Db: dbMock.DB}
	mock := dbMock.Sqlmock
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `messages` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	user, err := messageRepository.Insert(ctx, models.MessageCreate{
		FkUser:  1,
		Content: "Some dummy message",
	})
	assert := assert.New(t)
	assert.NoError(err)
	assert.NotNil(user)
}

func TestMessageRepository_Insert_Return_Error(t *testing.T) {
	ctx := context.Background()
	dbMock := test.NewMockDB()
	messageRepository := repositories.MessageRepository{Db: dbMock.DB}
	mock := dbMock.Sqlmock
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `messages` (.+)$").WillReturnError(errors.New("error"))
	mock.ExpectRollback()
	user, err := messageRepository.Insert(ctx, models.MessageCreate{
		FkUser:  1,
		Content: "Some dummy message",
	})
	assert := assert.New(t)
	assert.Error(err)
	assert.Nil(user)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

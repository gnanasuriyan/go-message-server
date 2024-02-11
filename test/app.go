package test

import (
	"github.com/gnanasuriyan/go-message-server/app/repositories"
	"github.com/gnanasuriyan/go-message-server/app/services"
	"github.com/gnanasuriyan/go-message-server/internal/db"
	"github.com/google/wire"
)

type TestAppOptions struct {
	// mock db
	AppDB db.IAppDB
	// repositories
	UserRepository repositories.IUserRepository
	// services
	UserService    services.IUserService
	MessageService services.IMessageService
}

type TestApp struct {
	*TestAppOptions
}

var NewTestApp = wire.NewSet(
	NewMockDB,
	wire.Bind(new(db.IAppDB), new(*MockAppDb)),
	wire.Struct(new(TestAppOptions), "*"),
	InitTestApp,
)

func InitTestApp(options *TestAppOptions) *TestApp {
	return &TestApp{
		TestAppOptions: options,
	}
}

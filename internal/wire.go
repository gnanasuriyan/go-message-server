//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/gnanasuriyan/go-message-server/app"
	"github.com/gnanasuriyan/go-message-server/app/repositories"
	"github.com/gnanasuriyan/go-message-server/app/services"
	"github.com/gnanasuriyan/go-message-server/internal/config"
	"github.com/gnanasuriyan/go-message-server/internal/db"
	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	repositories.NewUserRepository,
	repositories.NewMessageRepository,
)

var serviceSet = wire.NewSet(
	services.NewMessageService,
	services.NewUserService,
)

var configSet = wire.NewSet(
	config.GetConfig,
	wire.Bind(new(app.IAppConfig), new(*config.Config)),
	wire.Bind(new(db.IDatabaseConfig), new(*config.Config)),
)

var databaseSet = wire.NewSet(
	db.InitDatabase,
	wire.Bind(new(db.IAppDB), new(*db.AppDb)),
)

func GetServer() app.IServer {
	wire.Build(
		configSet,
		databaseSet,
		repositorySet,
		serviceSet,
		app.NewServer,
	)
	return &app.Server{}
}

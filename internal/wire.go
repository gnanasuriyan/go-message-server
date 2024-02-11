//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/gnanasuriyan/go-message-server/app"
	"github.com/gnanasuriyan/go-message-server/app/repositories"
	"github.com/gnanasuriyan/go-message-server/app/services"
	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	repositories.NewUserRepository,
)

var serviceSet = wire.NewSet(
	services.NewMessageService,
	services.NewUserService,
)

func GetServer() app.IServer {
	wire.Build(
		repositorySet,
		serviceSet,
		app.NewServer,
	)
	return &app.Server{}
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package internal

import (
	"github.com/gnanasuriyan/go-message-server/app"
	"github.com/gnanasuriyan/go-message-server/app/repositories"
	"github.com/gnanasuriyan/go-message-server/app/services"
	"github.com/gnanasuriyan/go-message-server/internal/config"
	"github.com/gnanasuriyan/go-message-server/internal/db"
	"github.com/google/wire"
)

// Injectors from wire.go:

func GetServer() app.IServer {
	appDb := db.InitDatabase()
	userRepository := &repositories.UserRepository{
		Db: appDb,
	}
	messageService := &services.MessageService{
		UserRepository: userRepository,
	}
	configConfig := config.GetConfig()
	server := &app.Server{
		MessageService: messageService,
		AppConfig:      configConfig,
	}
	return server
}

// wire.go:

var repositorySet = wire.NewSet(repositories.NewUserRepository)

var serviceSet = wire.NewSet(services.NewMessageService, services.NewUserService)

var configSet = wire.NewSet(config.GetConfig, wire.Bind(new(app.IAppConfig), new(*config.Config)), wire.Bind(new(db.IDatabaseConfig), new(*config.Config)))

var databaseSet = wire.NewSet(db.InitDatabase, wire.Bind(new(db.IAppDB), new(*db.AppDb)))

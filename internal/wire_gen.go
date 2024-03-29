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
	configConfig := config.GetConfig()
	appDb := db.InitDatabase()
	userRepository := &repositories.UserRepository{
		Db: appDb,
	}
	userService := &services.UserService{
		UserRepository: userRepository,
	}
	messageRepository := &repositories.MessageRepository{
		Db: appDb,
	}
	messageService := &services.MessageService{
		MessageRepository: messageRepository,
		UserRepository:    userRepository,
	}
	server := &app.Server{
		AppConfig:      configConfig,
		UserService:    userService,
		MessageService: messageService,
	}
	return server
}

// wire.go:

var repositorySet = wire.NewSet(repositories.NewUserRepository, repositories.NewMessageRepository)

var serviceSet = wire.NewSet(services.NewMessageService, services.NewUserService)

var configSet = wire.NewSet(config.GetConfig, wire.Bind(new(app.IAppConfig), new(*config.Config)), wire.Bind(new(db.IDatabaseConfig), new(*config.Config)))

var databaseSet = wire.NewSet(db.InitDatabase, wire.Bind(new(db.IAppDB), new(*db.AppDb)))

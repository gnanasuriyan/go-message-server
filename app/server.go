package app

import (
	"fmt"

	"github.com/gnanasuriyan/go-message-server/app/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type IAppConfig interface {
	GetPort() uint
}

type IServer interface {
	Start()
}

type Server struct {
	MessageService services.IMessageService
	AppConfig      IAppConfig
}

var NewServer = wire.NewSet(wire.Struct(new(Server), "*"), wire.Bind(new(IServer), new(*Server)))

func (a *Server) Start() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(fmt.Sprintf(":%d", a.AppConfig.GetPort()))
}

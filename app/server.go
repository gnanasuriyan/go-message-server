package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type IServer interface {
	Start()
}

type Server struct {
}

var NewServer = wire.NewSet(wire.Struct(new(Server), "*"), wire.Bind(new(IServer), new(*Server)))

func (a *Server) Start() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

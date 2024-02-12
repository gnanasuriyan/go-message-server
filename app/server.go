package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/gofiber/fiber/v2/middleware/csrf"

	"github.com/gofiber/fiber/v2/middleware/cors"

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
	AppConfig      IAppConfig
	UserService    services.IUserService
	MessageService services.IMessageService
}

var NewServer = wire.NewSet(wire.Struct(new(Server), "*"), wire.Bind(new(IServer), new(*Server)))

func (a *Server) Start() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(csrf.New())
	app.Use(requestid.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Pong!")
	})

	api := app.Group("/api/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})
	api.Post("/login", a.UserService.AuthenticateUser)
	api.Post("/signup", a.UserService.Signup)

	// protected routes
	api.Post("/message", a.MessageService.PostMessage)

	if err := app.Listen(fmt.Sprintf(":%d", a.AppConfig.GetPort())); err != nil {
		panic(err)
	}
}

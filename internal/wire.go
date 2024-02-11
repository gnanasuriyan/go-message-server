//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/gnanasuriyan/go-message-server/app"
	"github.com/google/wire"
)

func GetServer() app.IServer {
	wire.Build(app.NewServer)
	return &app.Server{}
}

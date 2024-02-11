package main

import "github.com/gnanasuriyan/go-message-server/internal"

func main() {
	server := internal.GetServer()
	server.Start()
}

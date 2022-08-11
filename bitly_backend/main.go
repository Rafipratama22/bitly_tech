package main

import (
	"bitly_backend/app/interface/container"
	"bitly_backend/app/interface/server"
)

func main() {
	server.StartServer(container.SetupContainer())
}

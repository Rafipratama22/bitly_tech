package server

import (
	"bitly_backend/app/interface/container"
	"bitly_backend/app/shared/config"
	"bitly_backend/app/transport"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer(container container.Container) {
	app := gin.Default()
	transport := transport.SetupTransport(&container)
	SetupRouter(transport, app, container.Middleware)

	fmt.Println(app.Run(config.Server.PORTHTTP))
}

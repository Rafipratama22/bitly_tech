package server

import (
	"bitly_backend/app/interface/middleware"
	"bitly_backend/app/transport"

	"github.com/gin-gonic/gin"
	
	"github.com/gin-contrib/cors"
)

func SetupRouter(tp *transport.Transport, app *gin.Engine, middleware middleware.Middlewared) {
	app.Use(cors.Default())
	v1 := app.Group("/api/bitly")
	{
		v1.POST("/register", tp.Transporter.Register)
		v1.POST("/login", tp.Transporter.Login)
		v1.PATCH("/click/:id", tp.Transporter.PatchGetClick)
		v1.POST("/create", tp.Transporter.CreateUrlWithoutUser)
		v1.Use(middleware.ValidatedToken)
		v1.GET("/link/:id", tp.Transporter.GetOneProduct)
		v1.GET("/user", tp.Transporter.GetInfoUser)
		v1.GET("/products", tp.Transporter.GetListProductByUserID)
		v1.POST("/link", tp.Transporter.CreateUrl)
		v1.PUT("/link/:id", tp.Transporter.UpdateUrl)
	}	
}
package middleware

import "github.com/gin-gonic/gin"

type Middlewared interface {
	ValidatedToken(ctx *gin.Context)
}

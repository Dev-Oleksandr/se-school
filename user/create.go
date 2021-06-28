package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/auth/pkg/auth"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase auth.UseCase) {
	h := newHandler(usecase)

	router.POST("/create", h.signUp)
	router.POST("/login", h.signIn)
}
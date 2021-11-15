package handler

import (
	"github.com/Presbyter/ddd/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (h *UserHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

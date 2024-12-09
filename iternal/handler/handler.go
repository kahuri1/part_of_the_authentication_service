package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/part_of_the_authentication_service/iternal/model"
)

type partAuthService interface {
	SignUpService(ctx context.Context, input model.UserSignUpInput) error
	AuthenticationService(request model.AuthenticationRequest) (model.Tokens, error)
	RefreshTokenService(token model.Tokens, ip string) (model.Tokens, error)
}

type Handler struct {
	service partAuthService
}

func Newhandler(service partAuthService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/singup", h.UserSingUpInput)
	r.POST("/auth", h.AuthenticationHandler)
	r.POST("/auth/refresh", h.RefreshTokenHandler)
	return r
}

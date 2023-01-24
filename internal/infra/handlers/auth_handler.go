package handlers

import (
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
)

type AuthHandler struct {
	UserGateway gateway.UserGateway
}

func NewAuthHandler(userGateway gateway.UserGateway) *AuthHandler {
	return &AuthHandler{
		UserGateway: userGateway,
	}
}

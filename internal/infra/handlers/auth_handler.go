package handlers

import (
	"github.com/go-chi/oauth"
	"net/http"
)

// todo: adding documentation
// todo: implement tests for this handler
type AuthHandler struct {
	BearerServer *oauth.BearerServer
}

func NewAuthHandler(bearerServer *oauth.BearerServer) *AuthHandler {
	return &AuthHandler{BearerServer: bearerServer}
}

func (h *AuthHandler) ClientCredentials(w http.ResponseWriter, r *http.Request) {
	h.BearerServer.ClientCredentials(w, r)
}

func (h *AuthHandler) UserCredentials(w http.ResponseWriter, r *http.Request) {
	h.BearerServer.UserCredentials(w, r)
}

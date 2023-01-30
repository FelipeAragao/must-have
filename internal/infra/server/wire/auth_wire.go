//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"
	"github.com/FelipeAragao/must-have/internal/infra/handlers"
	"github.com/FelipeAragao/must-have/internal/usecase/user/user_verifier"
	"github.com/go-chi/oauth"
	"github.com/google/wire"
)

func InitializeUserVerifier(db *sql.DB) *user_verifier.UserVerifier {
	wire.Build(
		setUserRepositoryDependency,
		user_verifier.NewUserVerifier,
	)
	return &user_verifier.UserVerifier{}
}

func InitializeAuthHandler(oauth *oauth.BearerServer) *handlers.AuthHandler {
	wire.Build(
		handlers.NewAuthHandler,
	)
	return &handlers.AuthHandler{}
}

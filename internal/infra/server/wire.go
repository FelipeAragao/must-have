//go:build wireinject
// +build wireinject

package server

import (
	"database/sql"
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"github.com/FelipeAragao/must-have/internal/infra/handlers"
	"github.com/FelipeAragao/must-have/internal/infra/repository"
	"github.com/FelipeAragao/must-have/internal/usecase/user/user_verifier"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(gateway.UserGateway), new(*repository.UserRepository)),
)

func InitializeUserVerifier(db *sql.DB) *user_verifier.UserVerifier {
	wire.Build(
		setRepositoryDependency,
		user_verifier.NewUserVerifier,
	)
	return &user_verifier.UserVerifier{}
}

func InitializeUserHandler(db *sql.DB) *handlers.UserHandler {
	wire.Build(
		setRepositoryDependency,
		handlers.NewUserHandler,
	)
	return &handlers.UserHandler{}
}

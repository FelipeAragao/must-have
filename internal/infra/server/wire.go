//go:build wireinject
// +build wireinject

package server

import (
	"database/sql"
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"github.com/FelipeAragao/must-have/internal/infra/database"
	"github.com/FelipeAragao/must-have/internal/usecase/user/user_verifier"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	database.NewUserRepository,
	wire.Bind(new(gateway.UserGateway), new(*database.UserRepository)),
)

func InitializeUserVerifier(db *sql.DB) *user_verifier.UserVerifier {
	wire.Build(
		setRepositoryDependency,
		user_verifier.NewUserVerifier,
	)
	return &user_verifier.UserVerifier{}
}

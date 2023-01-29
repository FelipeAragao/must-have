// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"database/sql"
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"github.com/FelipeAragao/must-have/internal/infra/handlers"
	"github.com/FelipeAragao/must-have/internal/infra/repository"
	"github.com/FelipeAragao/must-have/internal/usecase/user/change_password"
	"github.com/FelipeAragao/must-have/internal/usecase/user/create_user"
	"github.com/FelipeAragao/must-have/internal/usecase/user/find_by_id_user"
	"github.com/FelipeAragao/must-have/internal/usecase/user/update_user"
	"github.com/FelipeAragao/must-have/internal/usecase/user/user_verifier"
	"github.com/go-chi/oauth"
	"github.com/google/wire"
)

// Injectors from auth_wire.go:

func InitializeUserVerifier(db *sql.DB) *user_verifier.UserVerifier {
	userRepository := repository.NewUserRepository(db)
	userVerifier := user_verifier.NewUserVerifier(userRepository)
	return userVerifier
}

func InitializeAuthHandler(oauth2 *oauth.BearerServer) *handlers.AuthHandler {
	authHandler := handlers.NewAuthHandler(oauth2)
	return authHandler
}

// Injectors from user_wire.go:

// Initialize
func InitializeUserHandler(db *sql.DB) *handlers.UserHandler {
	userRepository := repository.NewUserRepository(db)
	createUserUseCase := create_user.NewCreateUserUseCase(userRepository)
	findByIDUseCase := find_by_id_user.NewFindByIDUseCase(userRepository)
	updateUserUseCase := update_user.NewUpdateUserUseCase(userRepository)
	changePasswordUserUseCase := change_password.NewChangePasswordUserUseCase(userRepository)
	userHandler := handlers.NewUserHandler(createUserUseCase, findByIDUseCase, updateUserUseCase, changePasswordUserUseCase)
	return userHandler
}

// user_wire.go:

// repository
var setUserRepositoryDependency = wire.NewSet(repository.NewUserRepository, wire.Bind(new(gateway.UserGateway), new(*repository.UserRepository)))

// usecase
var setCreateUseCaseDependency = wire.NewSet(create_user.NewCreateUserUseCase, wire.Bind(new(create_user.CreateUserUseCaseInterface), new(*create_user.CreateUserUseCase)))

var setUpdateUseCaseDependency = wire.NewSet(update_user.NewUpdateUserUseCase, wire.Bind(new(update_user.UpdateUserUseCaseInterface), new(*update_user.UpdateUserUseCase)))

var setFindByIDUseCaseDependency = wire.NewSet(find_by_id_user.NewFindByIDUseCase, wire.Bind(new(find_by_id_user.FindByIDUseCaseInterface), new(*find_by_id_user.FindByIDUseCase)))

var setChangePasswordUseCaseDependency = wire.NewSet(change_password.NewChangePasswordUserUseCase, wire.Bind(new(change_password.ChangePasswordUseCaseInterface), new(*change_password.ChangePasswordUserUseCase)))

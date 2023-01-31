//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"github.com/FelipeAragao/must-have/internal/infra/handlers"
	"github.com/FelipeAragao/must-have/internal/infra/repository"
	"github.com/FelipeAragao/must-have/internal/usecase/user/change_password"
	"github.com/FelipeAragao/must-have/internal/usecase/user/create_user"
	"github.com/FelipeAragao/must-have/internal/usecase/user/find_by_id"
	"github.com/FelipeAragao/must-have/internal/usecase/user/update_user"
	"github.com/google/wire"
)

// repository
var setUserRepositoryDependency = wire.NewSet(
	repository.NewUserRepository, wire.Bind(new(gateway.UserGateway), new(*repository.UserRepository)),
)

// usecase
var setCreateUseCaseDependency = wire.NewSet(
	create_user.NewCreateUserUseCase,
	wire.Bind(new(create_user.CreateUserUseCaseInterface), new(*create_user.CreateUserUseCase)),
)

var setUpdateUseCaseDependency = wire.NewSet(
	update_user.NewUpdateUserUseCase,
	wire.Bind(new(update_user.UpdateUserUseCaseInterface), new(*update_user.UpdateUserUseCase)),
)

var setFindByIDUseCaseDependency = wire.NewSet(
	find_by_id.NewFindByIDUseCase,
	wire.Bind(new(find_by_id.FindByIDUseCaseInterface), new(*find_by_id.FindByIDUseCase)),
)

var setChangePasswordUseCaseDependency = wire.NewSet(
	change_password.NewChangePasswordUserUseCase,
	wire.Bind(new(change_password.ChangePasswordUseCaseInterface), new(*change_password.ChangePasswordUserUseCase)),
)

// Initialize
func InitializeUserHandler(db *sql.DB) *handlers.UserHandler {
	wire.Build(
		setUserRepositoryDependency,
		setCreateUseCaseDependency,
		setUpdateUseCaseDependency,
		setFindByIDUseCaseDependency,
		setChangePasswordUseCaseDependency,
		handlers.NewUserHandler,
	)
	return &handlers.UserHandler{}
}

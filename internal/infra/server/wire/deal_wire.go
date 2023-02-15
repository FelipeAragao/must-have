//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"github.com/FelipeAragao/must-have/internal/infra/handlers"
	"github.com/FelipeAragao/must-have/internal/infra/repository"
	"github.com/FelipeAragao/must-have/internal/usecase/deal/create_deal"
	"github.com/FelipeAragao/must-have/internal/usecase/deal/find_by_id"
	"github.com/FelipeAragao/must-have/internal/usecase/deal/update_deal"
	"github.com/google/wire"
)

// repository
var setDealRepositoryDependency = wire.NewSet(
	repository.NewDealRepository, wire.Bind(new(gateway.DealGateway), new(*repository.DealRepository)),
)

// usecase
var setCreateDealUseCaseDependency = wire.NewSet(
	create_deal.NewCreateDealUseCase,
	wire.Bind(new(create_deal.CreateDealUseCaseInterface), new(*create_deal.CreateDealDealCase)),
)

var setUpdateDealUseCaseDependency = wire.NewSet(
	update_deal.NewUpdateDealUseCase,
	wire.Bind(new(update_deal.UpdateDealUseCaseInterface), new(*update_deal.UpdateDealDealCase)),
)

var setFindByIDDealUseCaseDependency = wire.NewSet(
	find_by_id.NewFindByIdUseCase,
	wire.Bind(new(find_by_id.FindByIDUseCaseInterface), new(*find_by_id.FindByIDUseCase)),
)

// Initialize
func InitializeDealHandler(db *sql.DB) *handlers.DealHandler {
	wire.Build(
		setDealRepositoryDependency,
		setUserRepositoryDependency,
		setCreateDealUseCaseDependency,
		setUpdateDealUseCaseDependency,
		setFindByIDDealUseCaseDependency,
		handlers.NewDealHandler,
	)
	return &handlers.DealHandler{}
}

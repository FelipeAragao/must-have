//go:generate mockery --name=DealGateway --filename=deal_gateway_mock.go --output=../../../test/mock/gateway --outpkg=mocks

package gateway

import (
	"context"
	"github.com/FelipeAragao/must-have/internal/domain/entity"
)

type DealGateway interface {
	Create(ctx context.Context, user *entity.Deal) error
	Update(ctx context.Context, user *entity.Deal) error
	FindByID(ctx context.Context, id string) (*entity.Deal, error)
	FindByUser(ctx context.Context, user_id string) (*entity.Deal, error)
}

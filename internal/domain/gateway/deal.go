//go:generate mockery --name=DealGateway --filename=deal_gateway_mock.go --output=../../../test/mock/gateway --outpkg=mocks

package gateway

import "github.com/FelipeAragao/must-have/internal/domain/entity"

type DealGateway interface {
	Create(user *entity.Deal) error
	Update(user *entity.Deal) error
	FindByID(id string) (*entity.Deal, error)
	FindByUser(user_id string) (*entity.Deal, error)
}

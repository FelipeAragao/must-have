//go:generate mockery --name=UserGateway --filename=user_gateway_mock.go --output=../../../test/mock/gateway --outpkg=mocks

package gateway

import "github.com/FelipeAragao/must-have/internal/domain/entity"

type UserGateway interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	FindByID(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByLogin(login string) (*entity.User, error)
}

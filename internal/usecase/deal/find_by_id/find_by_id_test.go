package find_by_id

import (
	"github.com/FelipeAragao/must-have/internal/domain/entity"
	entityPkg "github.com/FelipeAragao/must-have/pkg/entity"
	mocks "github.com/FelipeAragao/must-have/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	id := entityPkg.NewID()

	m := &mocks.DealGateway{}
	m.On("FindByID", mock.Anything).Return(&entity.Deal{
		ID:          id,
		Type:        entity.TROCA,
		Value:       20.0,
		Description: "Descrição do produto",
		TradeFor:    "Troco por um celular",
		Location: entity.LocationDeal{
			Lat:     -23.5506507,
			Lng:     -46.6333824,
			Address: "Rua Vergueiro, 3185",
			City:    "São Paulo",
			State:   "SP",
			ZipCode: 11111111,
		},
		Urgency: entity.Urgency{
			Type:      entity.BAIXA,
			LimitDate: time.Now(),
		},
		Photos:    []string{"https://www.google.com.br/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	uc := NewFindByIdUseCase(m)

	input := &DealInputDTO{
		ID: id.String(),
	}

	output, err := uc.Execute(nil, input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindByID", 1)
}

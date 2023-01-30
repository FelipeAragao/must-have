package update_deal

import (
	"github.com/FelipeAragao/must-have/internal/domain/entity"
	entityPkg "github.com/FelipeAragao/must-have/pkg/entity"
	mocks "github.com/FelipeAragao/must-have/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestUpdateUserUseCase_Execute(t *testing.T) {
	id := entityPkg.NewID()
	user_id := entityPkg.NewID()

	md := &mocks.DealGateway{}
	md.On("Update", mock.Anything).Return(nil)
	md.On("FindByID", mock.Anything).Return(&entity.Deal{
		ID:          id,
		UserID:      user_id,
		Type:        entity.DealType(entity.MEDIA),
		Value:       100000.00,
		Description: "Aliquip mollit aute aliqua ut. In nisi non est voluptate anim consequat veniam quis. Nostrud ullamco esse nostrud dolor elit eu esse nulla incididunt magna. Laboris eu quis laborum nostrud esse et consectetur exercitation. Duis minim irure proident sint occaecat cillum officia deserunt",
		TradeFor:    "Aliquip mollit aute aliqua ut. In nisi non est voluptate anim consequat veniam quis. Nostrud ullamco esse nostrud dolor elit eu esse nulla incididunt magna. Laboris eu quis laborum nostrud esse et consectetur exercitation. Duis minim irure proident sint occaecat cillum officia deserunt",
		Location: entity.LocationDeal{
			Lat:     -23.5506507,
			Lng:     -46.6333824,
			Address: "Rua Vergueiro, 3185",
			City:    "São Paulo",
			State:   "SP",
			ZipCode: 65000000,
		},
		Urgency: entity.Urgency{
			Type:      entity.UrgencyType(entity.TROCA),
			LimitDate: time.Now(),
		},
		Photos: []string{
			"https://picsum.photos/200/300",
			"https://picsum.photos/200/300",
			"https://picsum.photos/200/300",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	uc := NewUpdateDealUseCase(md)

	input := &DealInputDTO{
		UserID:      user_id.String(),
		Type:        int(entity.MEDIA),
		Value:       100000.00,
		Description: "Aliquip mollit aute aliqua ut. In nisi non est voluptate anim consequat veniam quis. Nostrud ullamco esse nostrud dolor elit eu esse nulla incididunt magna. Laboris eu quis laborum nostrud esse et consectetur exercitation. Duis minim irure proident sint occaecat cillum officia deserunt",
		TradeFor:    "Aliquip mollit aute aliqua ut. In nisi non est voluptate anim consequat veniam quis. Nostrud ullamco esse nostrud dolor elit eu esse nulla incididunt magna. Laboris eu quis laborum nostrud esse et consectetur exercitation. Duis minim irure proident sint occaecat cillum officia deserunt",
		Location: struct {
			Lat     float64 `json:"lat"`
			Lng     float64 `json:"lng"`
			Address string  `json:"address"`
			City    string  `json:"city"`
			State   string  `json:"state"`
			ZipCode int     `json:"zip_code"`
		}{
			Lat:     -67.868968,
			Lng:     -12.943469,
			Address: "460 Hamilton Avenue, Oberlin, Michigan, 2608",
			City:    "São Luis",
			State:   "MA",
			ZipCode: 11000000,
		},
		Urgency: struct {
			Type      int       `json:"type"`
			LimitDate time.Time `json:"limit_date"`
		}{
			Type:      int(entity.MEDIA),
			LimitDate: time.Now(),
		},
		Photos: []string{"https://www.google.com.br", "https://www.google.com.br"},
	}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.NotEmpty(t, output.CreatedAt)
	assert.NotEmpty(t, output.UpdatedAt)

	md.AssertExpectations(t)
	md.AssertNumberOfCalls(t, "Update", 1)
	md.AssertNumberOfCalls(t, "FindByID", 1)
}

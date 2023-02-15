package create_deal

import (
	"github.com/FelipeAragao/must-have/internal/domain/entity"
	entityPkg "github.com/FelipeAragao/must-have/pkg/entity"
	mocks "github.com/FelipeAragao/must-have/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	id := entityPkg.NewID()
	hash, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

	mu := &mocks.UserGateway{}
	mu.On("FindByID", mock.Anything).Return(&entity.User{
		ID:        id,
		Name:      "John Doe",
		Email:     "john.doe@vibbra.com.br",
		Login:     "john.doe",
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Location: entity.LocationUser{
			Lat:     -23.5506507,
			Lng:     -46.6333824,
			Address: "Rua Vergueiro, 3185",
			City:    "São Paulo",
			State:   "SP",
			ZipCode: 65000000,
		},
	}, nil)
	md := &mocks.DealGateway{}
	md.On("Create", mock.Anything).Return(nil)

	uc := NewCreateDealUseCase(md, mu)

	input := &DealInputDTO{
		UserID:      id.String(),
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

	output, err := uc.Execute(nil, input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.NotEmpty(t, output.CreatedAt)
	assert.NotEmpty(t, output.UpdatedAt)

	md.AssertExpectations(t)
	md.AssertNumberOfCalls(t, "Create", 1)

	mu.AssertExpectations(t)
	mu.AssertNumberOfCalls(t, "FindByID", 1)
}

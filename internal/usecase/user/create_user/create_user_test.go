package create_user

import (
	"errors"
	mocks "github.com/FelipeAragao/must-have/test/mock/gateway"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	m := &mocks.UserGateway{}
	m.On("Create", mock.Anything).Return(nil)
	m.On("FindByEmail", mock.Anything).Return(nil, errors.New("user not found"))
	m.On("FindByLogin", mock.Anything).Return(nil, errors.New("user not found"))

	uc := NewCreateUserUseCase(m)

	input := &UserInputDTO{
		Name:     "John Doe",
		Email:    "john.doe@vibbra.com.br",
		Login:    "john.doe",
		Password: "12345678",
		Location: struct {
			Lat     float64 `json:"lat"`
			Lng     float64 `json:"lng"`
			Address string  `json:"address"`
			City    string  `json:"city"`
			State   string  `json:"state"`
			ZipCode int     `json:"zip_code"`
		}{
			Lat:     -23.5506507,
			Lng:     -46.6333824,
			Address: "Rua Vergueiro, 3185",
			City:    "São Paulo",
			State:   "SP",
			ZipCode: 65000000,
		},
	}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.NotEmpty(t, output.CreatedAt)
	assert.NotEmpty(t, output.UpdatedAt)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Create", 1)
	m.AssertNumberOfCalls(t, "FindByEmail", 1)
	m.AssertNumberOfCalls(t, "FindByLogin", 1)
}

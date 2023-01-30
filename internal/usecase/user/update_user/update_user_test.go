package update_user

import (
	"errors"
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

	m := &mocks.UserGateway{}
	m.On("Update", mock.Anything).Return(nil)
	m.On("FindByEmail", mock.Anything).Return(nil, errors.New("user not found"))
	m.On("FindByLogin", mock.Anything).Return(nil, errors.New("user not found"))
	m.On("FindByID", mock.Anything).Return(&entity.User{
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

	uc := NewUpdateUserUseCase(m)

	input := &UserInputDTO{
		ID:    id.String(),
		Name:  "John Silva",
		Email: "john.silva@vibbra.com.br",
		Login: "john.silva",
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
	assert.Equal(t, output.ID, input.ID)
	assert.Equal(t, output.Name, input.Name)
	assert.Equal(t, output.Email, input.Email)

	assert.NotEmpty(t, output.CreatedAt)
	assert.NotEmpty(t, output.UpdatedAt)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Update", 1)
	m.AssertNumberOfCalls(t, "FindByEmail", 1)
	m.AssertNumberOfCalls(t, "FindByLogin", 1)
	m.AssertNumberOfCalls(t, "FindByID", 1)
}

package find_by_id

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

	m := &mocks.UserGateway{}
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

	uc := NewFindByIDUseCase(m)

	input := &UserInputDTO{
		ID: id.String(),
	}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "FindByID", 1)
}

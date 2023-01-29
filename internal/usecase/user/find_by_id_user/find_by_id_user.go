//go:generate mockery --name=FindByIDUseCaseInterface --filename=find_by_id_mock.go --output=../../../../test/mock/usecase --outpkg=mocks

package find_by_id_user

import (
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"github.com/FelipeAragao/must-have/pkg/validate"
	"time"
)

type UserInputDTO struct {
	ID string `json:"id"`
}

type UserOutputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Location struct {
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
		Address string  `json:"address"`
		City    string  `json:"city"`
		State   string  `json:"state"`
		ZipCode int     `json:"zip_code"`
	} `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindByIDUseCaseInterface interface {
	Execute(input *UserInputDTO) (*UserOutputDTO, error)
}

type FindByIDUseCase struct {
	UserGateway gateway.UserGateway
}

func NewFindByIDUseCase(userGateway gateway.UserGateway) *FindByIDUseCase {
	return &FindByIDUseCase{
		UserGateway: userGateway,
	}
}

func (uc *FindByIDUseCase) Execute(input *UserInputDTO) (*UserOutputDTO, error) {

	if input == nil || input.ID == "" {
		return nil, validate.ErrorMessage("id", "id is required")
	}

	user, err := uc.UserGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Login:     user.Login,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Location: struct {
			Lat     float64 `json:"lat"`
			Lng     float64 `json:"lng"`
			Address string  `json:"address"`
			City    string  `json:"city"`
			State   string  `json:"state"`
			ZipCode int     `json:"zip_code"`
		}{
			Lat:     user.Location.Lat,
			Lng:     user.Location.Lng,
			Address: user.Location.Address,
			City:    user.Location.City,
			State:   user.Location.State,
			ZipCode: user.Location.ZipCode,
		},
	}, nil
}

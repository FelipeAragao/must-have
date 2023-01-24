package find_by_id_user

import (
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"github.com/FelipeAragao/must-have/pkg/validate"
	"time"
)

type UserInputDTO struct {
	ID string
}

type UserOutputDTO struct {
	ID       string
	Name     string
	Email    string
	Login    string
	Password string
	Location struct {
		Lat     float64
		Lng     float64
		Address string
		City    string
		State   string
		ZipCode int
	}
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FindByIdUserUseCaseInterface interface {
	Execute(input *UserInputDTO) (*UserOutputDTO, error)
}

type FindByIdUserUseCase struct {
	UserGateway gateway.UserGateway
}

func NewFindByIdUserUseCase(userGateway gateway.UserGateway) *FindByIdUserUseCase {
	return &FindByIdUserUseCase{
		UserGateway: userGateway,
	}
}

func (uc *FindByIdUserUseCase) Execute(input *UserInputDTO) (*UserOutputDTO, error) {

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
			Lat     float64
			Lng     float64
			Address string
			City    string
			State   string
			ZipCode int
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

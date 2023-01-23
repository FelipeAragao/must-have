package update_user

import (
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"github.com/FelipeAragao/must-have/pkg/utils"
	"time"
)

type UserInputDTO struct {
	ID       string
	Name     string
	Email    string
	Login    string
	Location struct {
		Lat     float64
		Lng     float64
		Address string
		City    string
		State   string
		ZipCode int
	}
}

type UserOutputDTO struct {
	ID       string
	Name     string
	Email    string
	Login    string
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

type UpdateUserUseCaseInterface interface {
	Execute(input *UserInputDTO) (*UserOutputDTO, error)
}

type UpdateUserUseCase struct {
	UserGateway gateway.UserGateway
}

func NewUpdateUserUseCase(userGateway gateway.UserGateway) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		UserGateway: userGateway,
	}
}

func (uc *UpdateUserUseCase) Execute(input *UserInputDTO) (*UserOutputDTO, error) {

	user, err := uc.UserGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if input.Email != user.Email {
		_, err = uc.UserGateway.FindByEmail(input.Email)
		if err == nil {
			return nil, utils.ErrorMessage("email", "email already exists")
		}
	}

	if input.Login != user.Login {
		_, err = uc.UserGateway.FindByLogin(input.Login)
		if err == nil {
			return nil, utils.ErrorMessage("login", "login already exists")
		}
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Login = input.Login
	user.Location.Lat = input.Location.Lat
	user.Location.Lng = input.Location.Lng
	user.Location.Address = input.Location.Address
	user.Location.City = input.Location.City
	user.Location.State = input.Location.State
	user.Location.ZipCode = input.Location.ZipCode

	err = user.Modify()
	if err != nil {
		return nil, err
	}

	err = uc.UserGateway.Update(user)
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

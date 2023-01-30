//go:generate mockery --name=CreateUserUseCaseInterface --filename=create_user_usecase_mock.go --output=../../../../test/mock/usecase --outpkg=mocks

package create_user

import (
	"github.com/FelipeAragao/must-have/internal/domain/entity"
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"github.com/FelipeAragao/must-have/pkg/validate"
	"time"
)

type UserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Location struct {
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
		Address string  `json:"address"`
		City    string  `json:"city"`
		State   string  `json:"state"`
		ZipCode int     `json:"zip_code"`
	} `json:"location"`
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

type CreateUserUseCaseInterface interface {
	Execute(input *UserInputDTO) (*UserOutputDTO, error)
}

type CreateUserUseCase struct {
	UserGateway gateway.UserGateway
}

func NewCreateUserUseCase(userGateway gateway.UserGateway) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserGateway: userGateway,
	}
}

func (uc *CreateUserUseCase) Execute(input *UserInputDTO) (*UserOutputDTO, error) {

	user, err := entity.NewUser(input.Name, input.Email, input.Login, input.Password, entity.LocationUser{
		Lat:     input.Location.Lat,
		Lng:     input.Location.Lng,
		Address: input.Location.Address,
		City:    input.Location.City,
		State:   input.Location.State,
		ZipCode: input.Location.ZipCode,
	})

	if err != nil {
		return nil, err
	}

	_, err = uc.UserGateway.FindByEmail(user.Email)
	if err == nil {
		return nil, validate.ErrorMessage("email", "email already exists")
	}

	_, err = uc.UserGateway.FindByLogin(user.Login)
	if err == nil {
		return nil, validate.ErrorMessage("login", "login already exists")
	}

	err = uc.UserGateway.Create(user)
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

package change_password

import (
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"time"
)

type UserInputDTO struct {
	ID          string
	Password    string
	PasswordNew string
}

type UserOutputDTO struct {
	ID        string
	UpdatedAt time.Time
}

type ChangePasswordUserUseCaseInterface interface {
	Execute(input *UserInputDTO) (*UserOutputDTO, error)
}

type ChangePasswordUserUseCase struct {
	UserGateway gateway.UserGateway
}

func NewChangePasswordUserUseCase(userGateway gateway.UserGateway) *ChangePasswordUserUseCase {
	return &ChangePasswordUserUseCase{
		UserGateway: userGateway,
	}
}

func (uc *ChangePasswordUserUseCase) Execute(input *UserInputDTO) (*UserOutputDTO, error) {

	user, err := uc.UserGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	isValid := user.ValidatePassword(input.Password)

	if !isValid {
		return nil, err
	}

	err = user.ChangePassword(input.PasswordNew)
	if err != nil {
		return nil, err
	}

	err = uc.UserGateway.Update(user)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:        user.ID.String(),
		UpdatedAt: user.UpdatedAt,
	}, nil
}

package entity

import (
	"github.com/FelipeAragao/must-have/pkg/entity"
	"github.com/FelipeAragao/must-have/pkg/utils"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name" validate:"required,min=3,max=100"`
	Email    string    `json:"email" validate:"required,email"`
	Login    string    `json:"login" validate:"required,min=3,max=50"`
	Password string    `json:"-" validate:"required,min=8,max=15"`
	Location Location  `json:"location" validate:"required,dive"`
}

type Location struct {
	Lat     float64 `json:"lat" validate:"required,latitude"`
	Lng     float64 `json:"lng" validate:"required,longitude"`
	Address string  `json:"address" validate:"required,min=3,max=100"`
	City    string  `json:"city" validate:"required,min=3,max=100"`
	State   string  `json:"state" validate:"required,min=2,max=2"`
	ZipCode int     `json:"zip_code" validate:"required,gte=10000000,lte=99999999"`
}

func NewUser(name string, email string, login string, password string, location Location) (*User, error) {
	u := User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Login:    login,
		Password: password,
		Location: location,
	}

	err := u.validate()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (u User) Modify(name string, email string, login string, password string, location Location) (*User, error) {
	u.Name = name
	u.Email = email
	u.Login = login
	u.Password = password
	u.Location = location

	err := u.validate()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (u User) validate() error {
	err := utils.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

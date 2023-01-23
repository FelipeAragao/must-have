package entity

import (
	"github.com/FelipeAragao/must-have/pkg/entity"
	"github.com/FelipeAragao/must-have/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Email     string    `json:"email" validate:"required,email"`
	Login     string    `json:"login" validate:"required,min=3,max=50"`
	Password  string    `json:"-" validate:"required"`
	Location  Location  `json:"location" validate:"required,dive"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

	hash, err := generatePassword(password)
	if err != nil {
		return nil, err
	}

	u := User{
		ID:        entity.NewID(),
		Name:      name,
		Email:     email,
		Login:     login,
		Password:  hash,
		Location:  location,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.validate()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (u *User) Modify() error {
	err := u.validate()
	if err != nil {
		return err
	}

	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) ChangePassword(password string) error {
	hash, err := generatePassword(password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u User) validate() error {
	err := utils.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

func generatePassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", utils.ErrorMessage("password", "error generating password")
	}
	return string(hash), nil
}

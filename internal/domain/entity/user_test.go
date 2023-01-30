package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {

	name := "John Doe"
	email := "john.doe@vibbra.com.br"
	login := "john.doe"
	password := "12345678"

	location := LocationUser{}
	location.Lat = -23.5506507
	location.Lng = -46.6333824
	location.Address = "Rua Vergueiro, 3185"
	location.City = "São Paulo"
	location.State = "SP"
	location.ZipCode = 65000000

	user, err := NewUser(name, email, login, password, location)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, login, user.Login)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, location, user.Location)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
}

func TestNewUserWithErrors(t *testing.T) {

	name := ""
	email := "john.doevibbra.com.br"
	login := ""
	password := "123"

	location := LocationUser{}
	location.Lat = 0.0
	location.Lng = 0.0
	location.Address = ""
	location.City = ""
	location.State = ""
	location.ZipCode = 650

	user, err := NewUser(name, email, login, password, location)

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "{\"message\":\"Name is a required field\",\"field\":\"Name\"}|{\"message\":\"Email must be a valid email address\",\"field\":\"Email\"}|{\"message\":\"Login is a required field\",\"field\":\"Login\"}|{\"message\":\"Lat is a required field\",\"field\":\"Lat\"}|{\"message\":\"Lng is a required field\",\"field\":\"Lng\"}|{\"message\":\"Address is a required field\",\"field\":\"Address\"}|{\"message\":\"City is a required field\",\"field\":\"City\"}|{\"message\":\"State is a required field\",\"field\":\"State\"}|{\"message\":\"ZipCode must be 10,000,000 or greater\",\"field\":\"ZipCode\"}", err.Error())
}

func TestUser_Modify(t *testing.T) {

	name := "John Doe"
	email := "john.doe@vibbra.com.br"
	login := "john.doe"
	password := "12345678"

	location := LocationUser{}
	location.Lat = -23.5506507
	location.Lng = -46.6333824
	location.Address = "Rua Vergueiro, 3185"
	location.City = "São Paulo"
	location.State = "SP"
	location.ZipCode = 65000000

	user, _ := NewUser(name, email, login, password, location)

	user.Name = "John Silva"
	user.Email = "john.silva@vibbra.com.br"
	user.Password = "123456789"
	user.Login = "john.silva"
	user.Location.Lat = -23.5506507
	user.Location.Lng = -46.6333824
	user.Location.Address = "Rua Vergueiro, 3185"
	user.Location.City = "São Paulo"
	user.Location.State = "SP"
	user.Location.ZipCode = 65000000

	err := user.Modify()
	assert.Nil(t, err)
}

func TestUser_ModifyWithErrors(t *testing.T) {
	name := "John Doe"
	email := "john.doe@vibbra.com.br"
	login := "john.doe"
	password := "12345678"

	location := LocationUser{}
	location.Lat = -23.5506507
	location.Lng = -46.6333824
	location.Address = "Rua Vergueiro, 3185"
	location.City = "São Paulo"
	location.State = "SP"
	location.ZipCode = 65000000

	user, _ := NewUser(name, email, login, password, location)

	user.Name = ""
	user.Email = "john.silva.vibbra.com.br"
	user.Login = ""
	user.Location.Lat = -23.5506507
	user.Location.Lng = -46.6333824
	user.Location.Address = "Rua Vergueiro, 3185"
	user.Location.City = "São Paulo"
	user.Location.State = "SP"
	user.Location.ZipCode = 65000000

	err := user.Modify()
	assert.NotNil(t, err)
	assert.Equal(t, "{\"message\":\"Name is a required field\",\"field\":\"Name\"}|{\"message\":\"Email must be a valid email address\",\"field\":\"Email\"}|{\"message\":\"Login is a required field\",\"field\":\"Login\"}", err.Error())
}

func TestUser_ChangePassword(t *testing.T) {

	name := "John Doe"
	email := "john.doe@vibbra.com.br"
	login := "john.doe"
	password := "12345678"

	location := LocationUser{}
	location.Lat = -23.5506507
	location.Lng = -46.6333824
	location.Address = "Rua Vergueiro, 3185"
	location.City = "São Paulo"
	location.State = "SP"
	location.ZipCode = 65000000

	user, _ := NewUser(name, email, login, password, location)

	err := user.ChangePassword("123456789")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	name := "John Doe"
	email := "john.doe@vibbra.com.br"
	login := "john.doe"
	password := "12345678"

	location := LocationUser{}
	location.Lat = -23.5506507
	location.Lng = -46.6333824
	location.Address = "Rua Vergueiro, 3185"
	location.City = "São Paulo"
	location.State = "SP"
	location.ZipCode = 65000000

	user, _ := NewUser(name, email, login, password, location)
	assert.True(t, user.ValidatePassword("12345678"))
	assert.False(t, user.ValidatePassword("123456789"))
}

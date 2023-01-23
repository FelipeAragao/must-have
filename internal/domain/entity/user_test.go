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

	location := Location{}
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
	assert.Equal(t, password, user.Password)
	assert.Equal(t, location, user.Location)
}

func TestNewUserWithErrors(t *testing.T) {

	name := ""
	email := "john.doevibbra.com.br"
	login := ""
	password := "123"

	location := Location{}
	location.Lat = 0.0
	location.Lng = 0.0
	location.Address = ""
	location.City = ""
	location.State = ""
	location.ZipCode = 650

	user, err := NewUser(name, email, login, password, location)

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "{\"message\":\"Name is a required field\",\"field\":\"Name\"}|{\"message\":\"Email must be a valid email address\",\"field\":\"Email\"}|{\"message\":\"Login is a required field\",\"field\":\"Login\"}|{\"message\":\"Password must be at least 8 characters in length\",\"field\":\"Password\"}|{\"message\":\"Lat is a required field\",\"field\":\"Lat\"}|{\"message\":\"Lng is a required field\",\"field\":\"Lng\"}|{\"message\":\"Address is a required field\",\"field\":\"Address\"}|{\"message\":\"City is a required field\",\"field\":\"City\"}|{\"message\":\"State is a required field\",\"field\":\"State\"}|{\"message\":\"ZipCode must be 10,000,000 or greater\",\"field\":\"ZipCode\"}", err.Error())
}

func TestUser_Modify(t *testing.T) {

	name := "John Doe"
	email := "john.doe@vibbra.com.br"
	login := "john.doe"
	password := "12345678"

	location := Location{}
	location.Lat = -23.5506507
	location.Lng = -46.6333824
	location.Address = "Rua Vergueiro, 3185"
	location.City = "São Paulo"
	location.State = "SP"
	location.ZipCode = 65000000

	user, _ := NewUser(name, email, login, password, location)

	NameNew := "John Silva"
	emailNew := "john.silva@vibbra.com.br"
	passwordNew := "123456789"
	locationNew := Location{}
	locationNew.Lat = -23.5506507
	locationNew.Lng = -46.6333824
	locationNew.Address = "Rua Vergueiro, 3185"
	locationNew.City = "São Paulo"
	locationNew.State = "SP"
	locationNew.ZipCode = 65000000

	userNew, err := user.Modify(NameNew, emailNew, login, passwordNew, location)
	assert.Nil(t, err)
	assert.NotNil(t, userNew)
	assert.Equal(t, NameNew, userNew.Name)
	assert.Equal(t, emailNew, userNew.Email)
	assert.Equal(t, passwordNew, userNew.Password)
	assert.Equal(t, locationNew, userNew.Location)
}

func TestUser_ModifyWithErrors(t *testing.T) {

	name := "John Doe"
	email := "john.doe@vibbra.com.br"
	login := "john.doe"
	password := "12345678"

	location := Location{}
	location.Lat = -23.5506507
	location.Lng = -46.6333824
	location.Address = "Rua Vergueiro, 3185"
	location.City = "São Paulo"
	location.State = "SP"
	location.ZipCode = 65000000

	user, _ := NewUser(name, email, login, password, location)

	NameNew := ""
	emailNew := "john.silva.vibbra.com.br"
	passwordNew := "123456789"
	locationNew := Location{}
	locationNew.Lat = -23.5506507
	locationNew.Lng = -46.6333824
	locationNew.Address = "Rua Vergueiro, 3185"
	locationNew.City = "São Paulo"
	locationNew.State = "SP"
	locationNew.ZipCode = 65000000

	userNew, err := user.Modify(NameNew, emailNew, login, passwordNew, location)
	assert.NotNil(t, err)
	assert.Nil(t, userNew)
	assert.Equal(t, "{\"message\":\"Name is a required field\",\"field\":\"Name\"}|{\"message\":\"Email must be a valid email address\",\"field\":\"Email\"}", err.Error())
}

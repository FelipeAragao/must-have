package jwt

import (
	"fmt"
	"github.com/FelipeAragao/must-have/pkg/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateJWTToken(t *testing.T) {

	ID := entity.NewID()
	jwtSecret := "secret"
	jwtExpiresIn := 3600

	token, err := GenerateJWTToken(ID, jwtSecret, jwtExpiresIn)
	fmt.Println(token)
	assert.NotNil(t, token)
	assert.Nil(t, err)
}

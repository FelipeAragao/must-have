package jwt

import (
	"github.com/FelipeAragao/must-have/pkg/entity"
	"github.com/go-chi/jwtauth"
	"time"
)

func GenerateJWTToken(ID entity.ID, jwtSecret string, jwtExpiresIn int) (string, error) {

	jwt := jwtauth.New("HS256", []byte(jwtSecret), nil)

	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub": ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

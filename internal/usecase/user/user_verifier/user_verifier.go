package user_verifier

import (
	"errors"
	"github.com/FelipeAragao/must-have/internal/domain/gateway"
	"github.com/go-chi/oauth"
	"net/http"
)

type UserVerifier struct {
	gateway.UserGateway
}

func NewUserVerifier(userGateway gateway.UserGateway) *UserVerifier {
	return &UserVerifier{userGateway}
}

func (u *UserVerifier) ValidateUser(username, password, scope string, r *http.Request) error {
	user, err := u.UserGateway.FindByLogin(username)
	if err != nil {
		return err
	}

	if !user.ValidatePassword(password) {
		return errors.New("Not authorized")
	}

	return nil
}

func (u *UserVerifier) ValidateClient(clientID, clientSecret, scope string, r *http.Request) error {
	return nil
}

func (u *UserVerifier) AddClaims(tokenType oauth.TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {

	user, err := u.UserGateway.FindByLogin(credential)
	if err != nil {
		return nil, err
	}

	claims := make(map[string]string)
	claims["id"] = user.ID.String()
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["login"] = user.Login

	return claims, nil
}

func (u *UserVerifier) AddProperties(tokenType oauth.TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {

	user, err := u.UserGateway.FindByLogin(credential)
	if err != nil {
		return nil, err
	}

	props := make(map[string]string)
	props["user_id"] = user.ID.String()

	return props, nil
}

func (u *UserVerifier) ValidateTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	return nil
}

func (u *UserVerifier) StoreTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	return nil
}

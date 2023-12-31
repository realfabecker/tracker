package users

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"

	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

type CognitoAuthService struct {
	cognitoJwkUrl string
	jwtHandler    corpts.JwtHandler
}

func NewCognitoAuthService(
	cognitoJwkUrl string,
	jwtHandler corpts.JwtHandler,
) corpts.AuthService {
	return &CognitoAuthService{cognitoJwkUrl, jwtHandler}
}

func (u *CognitoAuthService) Verify(token string) (*jwt.RegisteredClaims, error) {
	c, err := u.jwtHandler.VerifyWithKeyURL(token, u.cognitoJwkUrl)
	if err != nil {
		return nil, err
	}
	if time.Now().Unix() > c.ExpiresAt.Time.Unix() {
		return nil, errors.New("token expired")
	}
	return c, nil
}

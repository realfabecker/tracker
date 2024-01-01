package users

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"

	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

type CognitoAuthService struct {
	jwtHandler corpts.JwtHandler
}

func NewCognitoAuthService(
	jwtHandler corpts.JwtHandler,
) corpts.AuthService {
	return &CognitoAuthService{jwtHandler}
}

func (u *CognitoAuthService) Verify(token string) (*jwt.RegisteredClaims, error) {
	c, err := u.jwtHandler.Decode(token)
	if err != nil {
		return nil, err
	}
	if time.Now().Unix() > c.ExpiresAt.Time.Unix() {
		return nil, errors.New("token expired")
	}
	return c, nil
}

package jwt2

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

type claims struct {
	jwt.RegisteredClaims
}

type JwtHandler struct{}

func NewJwtHandler() corpts.JwtHandler {
	return &JwtHandler{}
}

func (j *JwtHandler) Decode(token string) (*jwt.RegisteredClaims, error) {
	p := jwt.NewParser()
	t, _, err := p.ParseUnverified(token, &claims{})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	cl, ok := t.Claims.(*claims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return &cl.RegisteredClaims, nil
}

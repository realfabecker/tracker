package ports

import (
	"github.com/golang-jwt/jwt/v5"
)

type HttpHandler interface {
	Register() error
	Listen(port string) error
	GetApp() interface{}
}

type CacheHandler interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
}

type JwtHandler interface {
	Decode(token string) (*jwt.RegisteredClaims, error)
}

type AuthService interface {
	Verify(token string) (*jwt.RegisteredClaims, error)
}

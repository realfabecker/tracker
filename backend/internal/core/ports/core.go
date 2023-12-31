package ports

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
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
	FetchJWK(url string) (jwk.Set, error)
	VerifyWithKeyURL(token string, keyUrl string) (*jwt.RegisteredClaims, error)
	VerifyWithKeySet(t string, keySet jwk.Set) (*jwt.RegisteredClaims, error)
}

type AuthService interface {
	Verify(token string) (*jwt.RegisteredClaims, error)
}

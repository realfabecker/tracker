package sjwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

// Claims
type claims struct {
	jwt.RegisteredClaims
}

func VerifyWithKeySet(t string, keySet jwk.Set) (*jwt.RegisteredClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}
		key, _ := keySet.LookupKeyID(kid)
		if key == nil {
			return nil, fmt.Errorf("key %v not found", kid)
		}
		var raw interface{}
		return raw, key.Raw(&raw)
	}
	token, err := jwt.ParseWithClaims(t, &claims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	cl, ok := token.Claims.(*claims)
	if !token.Valid || !ok {
		return nil, errors.New("invalid token")
	}
	return &cl.RegisteredClaims, nil
}

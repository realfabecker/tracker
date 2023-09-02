package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/v2/jwk"

	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

// Claims
type claims struct {
	jwt.RegisteredClaims
}

// JwtHandler
type JwtHandler struct {
	cache corpts.CacheHandler
}

// NewJwtHandler
func NewJwtHandler(cache corpts.CacheHandler) corpts.JwtHandler {
	return &JwtHandler{cache}
}

// VerifyWithKeyURL
func (j *JwtHandler) VerifyWithKeyURL(token string, keyUrl string) (*jwt.RegisteredClaims, error) {
	keySet, err := j.FetchJWK(keyUrl)
	if err != nil {
		return nil, err
	}
	return j.VerifyWithKeySet(token, keySet)
}

// VerifyWithKeySet
func (j *JwtHandler) VerifyWithKeySet(t string, keySet jwk.Set) (*jwt.RegisteredClaims, error) {
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

// FetchJWK
func (j *JwtHandler) FetchJWK(url string) (jwk.Set, error) {
	d, err := j.cache.Get("key-set")
	if err != nil {
		return nil, err
	}
	if len(d) > 0 {
		return jwk.Parse(d)
	}

	client := http.Client{}
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("unable to obtain jwks from " + url)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	j.cache.Set("key-set", body)
	return jwk.Parse(body)
}

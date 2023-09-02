package services

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/golang-jwt/jwt/v5"

	corpts "github.com/realfabecker/wallet/internal/core/ports"
	usrdom "github.com/realfabecker/wallet/internal/users/domain"
	usrpts "github.com/realfabecker/wallet/internal/users/ports"
)

// AuthService
type CognitoAuthService struct {
	cognitoClient   *cognitoidentityprovider.Client
	cognitoClientId string
	cognitoJwkUrl   string
	jwtHandler      corpts.JwtHandler
}

// NewAuthService
func NewCognitoAuthService(
	cognitoClientId string,
	cognitoJwkUrl string,
	cognitoClient *cognitoidentityprovider.Client,
	jwtHandler corpts.JwtHandler,
) usrpts.AuthService {
	return &CognitoAuthService{cognitoClient, cognitoClientId, cognitoJwkUrl, jwtHandler}
}

// Login
func (u *CognitoAuthService) Login(email string, password string) (*usrdom.UserToken, error) {
	auth, err := u.cognitoClient.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		ClientId: aws.String(u.cognitoClientId),
		AuthParameters: map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		},
	})
	if err != nil {
		return nil, err
	}
	return &usrdom.UserToken{
		AccesToken:   *auth.AuthenticationResult.AccessToken,
		RefreshToken: *auth.AuthenticationResult.RefreshToken,
	}, nil
}

//Verify
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

package ports

import (
	"github.com/golang-jwt/jwt/v5"

	usrdom "github.com/realfabecker/wallet/internal/users/domain"
)

// AuthService
type AuthService interface {
	Login(email string, password string) (*usrdom.UserToken, error)
	Change(email string, password string, session string) (*usrdom.UserToken, error)
	Verify(token string) (*jwt.RegisteredClaims, error)
}

// UserRepository
type UserRepository interface {
	GetUserByEmail(email string) (*usrdom.User, error)
}

// UserService
type UserService interface {
	GetUserByEmail(email string) (*usrdom.User, error)
}

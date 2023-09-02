package services

import (
	usrdom "github.com/realfabecker/wallet/internal/users/domain"
	usrpts "github.com/realfabecker/wallet/internal/users/ports"
)

// UserService
type UserService struct {
	UserRepository usrpts.UserRepository
}

// NewUserService
func NewUserService(r usrpts.UserRepository) usrpts.UserService {
	return &UserService{UserRepository: r}
}

// GetUserByEmail
func (s *UserService) GetUserByEmail(email string) (*usrdom.User, error) {
	return s.UserRepository.GetUserByEmail(email)
}

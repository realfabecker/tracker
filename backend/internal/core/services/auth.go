package services

import (
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

// UserService
type UserService struct {
	UserRepository corpts.UserRepository
}

// NewUserService
func NewUserService(r corpts.UserRepository) corpts.UserService {
	return &UserService{UserRepository: r}
}

// GetUserByEmail
func (s *UserService) GetUserByEmail(email string) (*cordom.User, error) {
	return s.UserRepository.GetUserByEmail(email)
}

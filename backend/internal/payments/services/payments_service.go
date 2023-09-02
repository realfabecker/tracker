package services

import (
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	paydom "github.com/realfabecker/wallet/internal/payments/domain"
	paypts "github.com/realfabecker/wallet/internal/payments/ports"
)

type WalletService struct {
	WalletRepository paypts.WalletRepository
}

func NewWalletService(r paypts.WalletRepository) paypts.WalletService {
	return &WalletService{WalletRepository: r}
}

func (s *WalletService) ListPayments(email string, q paydom.PaymentPagedDTOQuery) (*cordom.PagedDTO[paydom.Payment], error) {
	return s.WalletRepository.ListPayments(email, q)
}

func (s *WalletService) CreatePayment(p *paydom.Payment) (*paydom.Payment, error) {
	return s.WalletRepository.CreatePayment(p)
}

func (s *WalletService) PutPayment(p *paydom.Payment) (*paydom.Payment, error) {
	return s.WalletRepository.PutPayment(p)
}

func (s *WalletService) GetPaymentById(user string, payment string) (*paydom.Payment, error) {
	return s.WalletRepository.GetPaymentById(user, payment)
}

func (s *WalletService) DeletePayment(user string, payment string) error {
	return s.WalletRepository.DeletePayment(user, payment)
}

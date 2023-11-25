package services

import (
	"errors"

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

func (s *WalletService) CreateTransactionDetail(user string, d *paydom.TransactionDetail) (*paydom.TransactionDetail, error) {
	if p, err := s.GetPaymentById(user, d.TransactionId); err != nil {
		return nil, err
	} else if p == nil {
		return nil, errors.New("transaction does not exists")
	}
	return s.WalletRepository.CreateTransactionDetail(d)
}

func (s *WalletService) ListTransactionDetails(payment string, q cordom.PagedDTOQuery) (*cordom.PagedDTO[paydom.TransactionDetail], error) {
	return s.WalletRepository.ListTransactionDetails(payment, q)
}

func (s *WalletService) GetTransactionDetail(transactionId string, detailId string) (*paydom.TransactionDetail, error) {
	return s.WalletRepository.GetTransactionDetail(transactionId, detailId)
}

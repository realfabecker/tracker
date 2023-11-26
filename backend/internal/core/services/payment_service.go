package services

import (
	"errors"

	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

type WalletService struct {
	WalletRepository corpts.WalletRepository
}

func NewWalletService(r corpts.WalletRepository) corpts.WalletService {
	return &WalletService{WalletRepository: r}
}

func (s *WalletService) ListPayments(email string, q cordom.PaymentPagedDTOQuery) (*cordom.PagedDTO[cordom.Payment], error) {
	return s.WalletRepository.ListPayments(email, q)
}

func (s *WalletService) CreatePayment(p *cordom.Payment) (*cordom.Payment, error) {
	return s.WalletRepository.CreatePayment(p)
}

func (s *WalletService) PutPayment(p *cordom.Payment) (*cordom.Payment, error) {
	return s.WalletRepository.PutPayment(p)
}

func (s *WalletService) GetPaymentById(user string, payment string) (*cordom.Payment, error) {
	return s.WalletRepository.GetPaymentById(user, payment)
}

func (s *WalletService) DeletePayment(user string, payment string) error {
	return s.WalletRepository.DeletePayment(user, payment)
}

func (s *WalletService) CreateTransactionDetail(user string, d *cordom.TransactionDetail) (*cordom.TransactionDetail, error) {
	if p, err := s.GetPaymentById(user, d.TransactionId); err != nil {
		return nil, err
	} else if p == nil {
		return nil, errors.New("transaction does not exists")
	}
	return s.WalletRepository.CreateTransactionDetail(d)
}

func (s *WalletService) ListTransactionDetails(payment string, q cordom.PagedDTOQuery) (*cordom.PagedDTO[cordom.TransactionDetail], error) {
	return s.WalletRepository.ListTransactionDetails(payment, q)
}

func (s *WalletService) GetTransactionDetail(transactionId string, detailId string) (*cordom.TransactionDetail, error) {
	return s.WalletRepository.GetTransactionDetail(transactionId, detailId)
}

package ports

import (
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	paydom "github.com/realfabecker/wallet/internal/payments/domain"
)

type WalletRepository interface {
	ListPayments(email string, q paydom.PaymentPagedDTOQuery) (*cordom.PagedDTO[paydom.Payment], error)
	CreatePayment(p *paydom.Payment) (*paydom.Payment, error)
	PutPayment(p *paydom.Payment) (*paydom.Payment, error)
	GetPaymentById(user string, payment string) (*paydom.Payment, error)
	DeletePayment(user string, payment string) error
}

type WalletService interface {
	ListPayments(email string, q paydom.PaymentPagedDTOQuery) (*cordom.PagedDTO[paydom.Payment], error)
	CreatePayment(p *paydom.Payment) (*paydom.Payment, error)
	PutPayment(p *paydom.Payment) (*paydom.Payment, error)
	GetPaymentById(user string, payment string) (*paydom.Payment, error)
	DeletePayment(user string, payment string) error
}

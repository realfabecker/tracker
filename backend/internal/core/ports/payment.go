package ports

import (
	cordom "github.com/realfabecker/wallet/internal/core/domain"
)

type WalletRepository interface {
	ListPayments(email string, q cordom.PaymentPagedDTOQuery) (*cordom.PagedDTO[cordom.Payment], error)
	CreatePayment(p *cordom.Payment) (*cordom.Payment, error)
	PutPayment(p *cordom.Payment) (*cordom.Payment, error)
	GetPaymentById(user string, payment string) (*cordom.Payment, error)
	DeletePayment(user string, payment string) error
	CreateTransactionDetail(p *cordom.TransactionDetail) (*cordom.TransactionDetail, error)
	ListTransactionDetails(payments string, q cordom.PagedDTOQuery) (*cordom.PagedDTO[cordom.TransactionDetail], error)
	GetTransactionDetail(transactionId string, detailId string) (*cordom.TransactionDetail, error)
}

type WalletService interface {
	ListPayments(email string, q cordom.PaymentPagedDTOQuery) (*cordom.PagedDTO[cordom.Payment], error)
	CreatePayment(p *cordom.Payment) (*cordom.Payment, error)
	PutPayment(p *cordom.Payment) (*cordom.Payment, error)
	GetPaymentById(user string, payment string) (*cordom.Payment, error)
	DeletePayment(user string, payment string) error
	CreateTransactionDetail(user string, p *cordom.TransactionDetail) (*cordom.TransactionDetail, error)
	ListTransactionDetails(payments string, q cordom.PagedDTOQuery) (*cordom.PagedDTO[cordom.TransactionDetail], error)
	GetTransactionDetail(transactionId string, detailId string) (*cordom.TransactionDetail, error)
}

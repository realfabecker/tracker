package ports

import (
	cordom "github.com/realfabecker/wallet/internal/core/domain"
)

type TransactionRepository interface {
	ListTransactions(email string, q cordom.TransactionPagedDTOQuery) (*cordom.PagedDTO[cordom.Transaction], error)
	CreateTransaction(p *cordom.Transaction) (*cordom.Transaction, error)
	PutTransaction(p *cordom.Transaction) (*cordom.Transaction, error)
	GetTransactionById(user string, transaction string) (*cordom.Transaction, error)
	DeleteTransaction(user string, transaction string) error
	CreateTransactionDetail(p *cordom.TransactionDetail) (*cordom.TransactionDetail, error)
	ListTransactionDetails(transactions string, q cordom.PagedDTOQuery) (*cordom.PagedDTO[cordom.TransactionDetail], error)
	GetTransactionDetail(transactionId string, detailId string) (*cordom.TransactionDetail, error)
	DeleteTransactionDetail(transactionId string, detailId string) error
	PutTransactionDetail(p *cordom.TransactionDetail) (*cordom.TransactionDetail, error)
}

type TransactionService interface {
	ListTransactions(email string, q cordom.TransactionPagedDTOQuery) (*cordom.PagedDTO[cordom.Transaction], error)
	CreateTransaction(p *cordom.Transaction) (*cordom.Transaction, error)
	PutTransaction(p *cordom.Transaction) (*cordom.Transaction, error)
	GetTransactionById(user string, transaction string) (*cordom.Transaction, error)
	DeleteTransaction(user string, transaction string) error
	CreateTransactionDetail(user string, p *cordom.TransactionDetail) (*cordom.TransactionDetail, error)
	ListTransactionDetails(transactions string, q cordom.PagedDTOQuery) (*cordom.PagedDTO[cordom.TransactionDetail], error)
	GetTransactionDetail(transactionId string, detailId string) (*cordom.TransactionDetail, error)
	DeleteTransactionDetail(transactionId string, detailId string) error
	PutTransactionDetail(user string, p *cordom.TransactionDetail) (*cordom.TransactionDetail, error)
}

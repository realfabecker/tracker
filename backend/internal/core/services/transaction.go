package services

import (
	"errors"

	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

type TransactionService struct {
	TransactionRepository corpts.TransactionRepository
}

func NewTransactionService(r corpts.TransactionRepository) corpts.TransactionService {
	return &TransactionService{TransactionRepository: r}
}

func (s *TransactionService) ListTransactions(email string, q cordom.TransactionPagedDTOQuery) (*cordom.PagedDTO[cordom.Transaction], error) {
	return s.TransactionRepository.ListTransactions(email, q)
}

func (s *TransactionService) CreateTransaction(p *cordom.Transaction) (*cordom.Transaction, error) {
	return s.TransactionRepository.CreateTransaction(p)
}

func (s *TransactionService) PutTransaction(p *cordom.Transaction) (*cordom.Transaction, error) {
	return s.TransactionRepository.PutTransaction(p)
}

func (s *TransactionService) GetTransactionById(user string, transaction string) (*cordom.Transaction, error) {
	return s.TransactionRepository.GetTransactionById(user, transaction)
}

func (s *TransactionService) DeleteTransaction(user string, transaction string) error {
	return s.TransactionRepository.DeleteTransaction(user, transaction)
}

func (s *TransactionService) CreateTransactionDetail(user string, d *cordom.TransactionDetail) (*cordom.TransactionDetail, error) {
	if p, err := s.GetTransactionById(user, d.TransactionId); err != nil {
		return nil, err
	} else if p == nil {
		return nil, errors.New("transaction does not exists")
	}
	return s.TransactionRepository.CreateTransactionDetail(d)
}

func (s *TransactionService) ListTransactionDetails(transaction string, q cordom.PagedDTOQuery) (*cordom.PagedDTO[cordom.TransactionDetail], error) {
	return s.TransactionRepository.ListTransactionDetails(transaction, q)
}

func (s *TransactionService) GetTransactionDetail(transactionId string, detailId string) (*cordom.TransactionDetail, error) {
	return s.TransactionRepository.GetTransactionDetail(transactionId, detailId)
}

func (s *TransactionService) DeleteTransactionDetail(transactionId string, detailId string) error {
	return s.TransactionRepository.DeleteTransactionDetail(transactionId, detailId)
}

func (s *TransactionService) PutTransactionDetail(user string, d *cordom.TransactionDetail) (*cordom.TransactionDetail, error) {
	if p, err := s.GetTransactionById(user, d.TransactionId); err != nil {
		return nil, err
	} else if p == nil {
		return nil, errors.New("transaction does not exists")
	}
	return s.TransactionRepository.PutTransactionDetail(d)
}

package repositories

import (
	cordom "github.com/realfabecker/wallet/internal/core/domain"
)

// transaction model info
type transaction struct {
	*cordom.Transaction
	PK     string `dynamodbav:"PK" json:"-"`
	SK     string `dynamodbav:"SK" json:"-"`
	GSI1PK string `dynamodbav:"GSI1_PK" json:"-"`
	GSI1SK string `dynamodbav:"GSI1_SK" json:"-"`
}

// transactionDetail model info
type transactionDetail struct {
	*cordom.TransactionDetail
	PK string `dynamodbav:"PK" json:"-"`
	SK string `dynamodbav:"SK" json:"-"`
}

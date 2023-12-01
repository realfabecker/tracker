package repositories

import (
	cordom "github.com/realfabecker/wallet/internal/core/domain"
)






// payment model info
type payment struct {
	*cordom.Payment
	PK     string `dynamodbav:"PK" json:"-"`
	SK     string `dynamodbav:"SK" json:"-"`
	GSI1PK string `dynamodbav:"GSI1_PK" json:"-"`
	GSI1SK string `dynamodbav:"GSI1_SK" json:"-"`
}

type transactionDetail struct {
	*cordom.TransactionDetail
	PK string `dynamodbav:"PK" json:"-"`
	SK string `dynamodbav:"SK" json:"-"`
}

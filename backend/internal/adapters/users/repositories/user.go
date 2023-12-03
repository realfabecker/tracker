package repositories

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

// TransactionRepository
type UserRepository struct {
	db    *dynamodb.Client
	table string
	app   string
}

// NewUserRepository
func NewUserRepository(db *dynamodb.Client, table string, app string) (corpts.UserRepository, error) {
	return &UserRepository{db, table, app}, nil
}

// GetUserByEmail
func (u *UserRepository) GetUserByEmail(email string) (*cordom.User, error) {
	d, err := u.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#USER#" + email,
			},
			"SK": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#USER#" + email,
			},
		},
		TableName: aws.String(u.table),
	})
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, errors.New("user not found")
	}
	out := &cordom.User{}
	if err := attributevalue.UnmarshalMap(d.Item, out); err != nil {
		return nil, err
	}
	return out, err
}

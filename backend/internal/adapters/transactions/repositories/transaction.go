package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/realfabecker/wallet/internal/adapters/common/dynamo"
	"github.com/realfabecker/wallet/internal/adapters/common/validator"
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

// TransactionRepository
type WalletDynamoDbRepository struct {
	db    *dynamodb.Client
	table string
	app   string
}

// NewWalletDynamoDBRepository
func NewWalletDynamoDBRepository(db *dynamodb.Client, table string, app string) (corpts.TransactionRepository, error) {
	return &WalletDynamoDbRepository{db, table, app}, nil
}

// ListUserTransactions
func (u *WalletDynamoDbRepository) ListTransactions(user string, q cordom.TransactionPagedDTOQuery) (*cordom.PagedDTO[cordom.Transaction], error) {
	cipher := fmt.Sprintf("%s%d", user, q.Limit)
	k, err := dynamo.DecodePageToken(q.PageToken, cipher)
	if err != nil {
		return nil, err
	}

	var out *dynamodb.QueryOutput
	out, err = u.db.Query(context.TODO(), &dynamodb.QueryInput{
		KeyConditionExpression: aws.String("PK = :v and begins_with(SK, :x)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#USER#" + user,
			},
			":x": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#MOVT#" + q.GetDueDate(),
			},
		},
		ScanIndexForward:  aws.Bool(false),
		TableName:         aws.String(u.table),
		Limit:             &q.Limit,
		ExclusiveStartKey: k,
	})

	if err != nil {
		return nil, err
	}

	var lst []cordom.Transaction
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &lst); err != nil {
		return nil, err
	}

	dto := cordom.PagedDTO[cordom.Transaction]{}
	dto.PageCount = out.ScannedCount
	dto.Items = lst
	dto.HasMore = out.LastEvaluatedKey != nil

	if out.LastEvaluatedKey != nil {
		if dto.PageToken, err = dynamo.EncodePageToken(
			out.LastEvaluatedKey,
			cipher,
		); err != nil {
			return nil, err
		}
	}
	return &dto, nil
}

// CreateTransaction
func (u *WalletDynamoDbRepository) CreateTransaction(p *cordom.Transaction) (*cordom.Transaction, error) {
	dueDateTime, err := validator.DateParse(p.DueDate)
	if err != nil {
		return nil, err
	}
	pd := transaction{Transaction: p}
	pd.TransactionId = dueDateTime.Format("20060102") + validator.NewULID(time.Now())
	pd.CreatedAt = time.Now().Format("2006-01-02T15:04:05-07:00")

	pd.PK = "APP#" + u.app + "#USER#" + p.UserId
	pd.SK = "APP#" + u.app + "#MOVT#" + p.TransactionId

	pd.GSI1PK = "APP#" + u.app + "#MOVT_STATUS#" + p.Status.String()
	pd.GSI1SK = "APP#" + u.app + "#MOVT_DUEDATE#" + dueDateTime.Format("20060102")

	avs, err := attributevalue.MarshalMap(pd)
	if err != nil {
		return nil, err
	}

	if _, err := u.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(u.table),
		Item:      avs,
	}); err != nil {
		return nil, err
	}

	return pd.Transaction, nil
}

// GetTransactionById
func (u *WalletDynamoDbRepository) GetTransactionById(user string, transaction string) (*cordom.Transaction, error) {
	out, err := u.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#USER#" + user,
			},
			"SK": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#MOVT#" + transaction,
			},
		},
		TableName: aws.String(u.table),
	})

	if err != nil {
		return nil, err
	}
	if out.Item == nil {
		return nil, nil
	}
	var dto cordom.Transaction
	if err := attributevalue.UnmarshalMap(out.Item, &dto); err != nil {
		return nil, err
	}
	return &dto, nil
}

// DeleteTransaction
func (u *WalletDynamoDbRepository) DeleteTransaction(user string, transaction string) error {
	_, err := u.db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#USER#" + user,
			},
			"SK": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#MOVT#" + transaction,
			},
		},
		TableName: aws.String(u.table),
	})
	return err
}

// PutTransaction
func (u *WalletDynamoDbRepository) PutTransaction(p *cordom.Transaction) (*cordom.Transaction, error) {
	r, err := u.GetTransactionById(p.UserId, p.TransactionId)
	if err != nil {
		return nil, err
	}

	if r == nil {
		return u.CreateTransaction(p)
	}

	dueDateTime, err := validator.DateParse(p.DueDate)
	if err != nil {
		return nil, err
	}

	pd := transaction{Transaction: p}
	pd.PK = "APP#" + u.app + "#USER#" + p.UserId
	pd.SK = "APP#" + u.app + "#MOVT#" + p.TransactionId

	pd.GSI1PK = "APP#" + u.app + "#MOVT_STATUS#" + p.Status.String()
	pd.GSI1SK = "APP#" + u.app + "#MOVT_DUEDATE#" + dueDateTime.Format("20060102")
	avs, err := attributevalue.MarshalMap(pd)
	if err != nil {
		return nil, err
	}

	if _, err := u.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(u.table),
		Item:      avs,
	}); err != nil {
		return nil, err
	}

	return p, nil
}

// CreateTransactionDetail
func (u *WalletDynamoDbRepository) CreateTransactionDetail(i *cordom.TransactionDetail) (*cordom.TransactionDetail, error) {
	td := transactionDetail{TransactionDetail: i}
	td.DetailId = validator.NewULID(time.Now())
	td.CreatedAt = time.Now().Format("2006-01-02T15:04:05-07:00")

	td.PK = "APP#" + u.app + "#MOVT#" + i.TransactionId
	td.SK = "APP#" + u.app + "#MOVT_DETAIL#" + i.DetailId

	avs, err := attributevalue.MarshalMap(td)
	if err != nil {
		return nil, err
	}

	if _, err := u.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(u.table),
		Item:      avs,
	}); err != nil {
		return nil, err
	}

	return td.TransactionDetail, nil
}

// ListTransactionDetails
func (u *WalletDynamoDbRepository) ListTransactionDetails(transaction string, q cordom.PagedDTOQuery) (*cordom.PagedDTO[cordom.TransactionDetail], error) {
	cipher := fmt.Sprintf("%s%d", transaction, q.Limit)
	k, err := dynamo.DecodePageToken(q.PageToken, cipher)
	if err != nil {
		return nil, err
	}

	var out *dynamodb.QueryOutput
	out, err = u.db.Query(context.TODO(), &dynamodb.QueryInput{
		KeyConditionExpression: aws.String("PK = :v and begins_with(SK, :x)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v": &types.AttributeValueMemberS{
				Value: "APP#wallet#MOVT#" + transaction,
			},
			":x": &types.AttributeValueMemberS{
				Value: "APP#wallet#MOVT_DETAIL",
			},
		},
		ScanIndexForward:  aws.Bool(false),
		TableName:         aws.String(u.table),
		Limit:             &q.Limit,
		ExclusiveStartKey: k,
	})

	if err != nil {
		return nil, err
	}

	var lst []cordom.TransactionDetail
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &lst); err != nil {
		return nil, err
	}

	dto := cordom.PagedDTO[cordom.TransactionDetail]{}
	dto.PageCount = out.ScannedCount
	dto.Items = lst
	dto.HasMore = out.LastEvaluatedKey != nil

	if out.LastEvaluatedKey != nil {
		if dto.PageToken, err = dynamo.EncodePageToken(
			out.LastEvaluatedKey,
			cipher,
		); err != nil {
			return nil, err
		}
	}
	return &dto, nil
}

// GetTrasactionDetail
func (u *WalletDynamoDbRepository) GetTransactionDetail(transactionId string, detailId string) (*cordom.TransactionDetail, error) {
	out, err := u.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: "APP#wallet#MOVT#" + transactionId,
			},
			"SK": &types.AttributeValueMemberS{
				Value: "APP#wallet#MOVT_DETAIL#" + detailId,
			},
		},
		TableName: aws.String(u.table),
	})
	if err != nil {
		return nil, err
	}
	if out.Item == nil {
		return nil, nil
	}
	var dto cordom.TransactionDetail
	if err := attributevalue.UnmarshalMap(out.Item, &dto); err != nil {
		return nil, err
	}
	return &dto, nil
}

// GetTrasactionDetail
func (u *WalletDynamoDbRepository) DeleteTransactionDetail(transactionId string, detailId string) error {
	_, err := u.db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: "APP#wallet#MOVT#" + transactionId,
			},
			"SK": &types.AttributeValueMemberS{
				Value: "APP#wallet#MOVT_DETAIL#" + detailId,
			},
		},
		TableName: aws.String(u.table),
	})
	return err
}

// CreateTransactionDetail
func (u *WalletDynamoDbRepository) PutTransactionDetail(i *cordom.TransactionDetail) (*cordom.TransactionDetail, error) {
	r, err := u.GetTransactionDetail(i.TransactionId, i.DetailId)
	if err != nil {
		return nil, err
	}
	if r == nil {
		return u.CreateTransactionDetail(i)
	}

	td := transactionDetail{TransactionDetail: i}
	td.PK = "APP#" + u.app + "#MOVT#" + i.TransactionId
	td.SK = "APP#" + u.app + "#MOVT_DETAIL#" + i.DetailId

	avs, err := attributevalue.MarshalMap(td)
	if err != nil {
		return nil, err
	}

	if _, err := u.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(u.table),
		Item:      avs,
	}); err != nil {
		return nil, err
	}

	return td.TransactionDetail, nil
}

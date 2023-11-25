package dynamodb

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	"github.com/realfabecker/wallet/internal/core/lib/dynamo"
	"github.com/realfabecker/wallet/internal/core/lib/validator"
	paydom "github.com/realfabecker/wallet/internal/payments/domain"
	paypts "github.com/realfabecker/wallet/internal/payments/ports"
)

//WalletRepository
type WalletDynamoDbRepository struct {
	db    *dynamodb.Client
	table string
	app   string
}

// NewWalletDynamoDBRepository
func NewWalletDynamoDBRepository(db *dynamodb.Client, table string, app string) (paypts.WalletRepository, error) {
	return &WalletDynamoDbRepository{db, table, app}, nil
}

//ListUserPayments
func (u *WalletDynamoDbRepository) ListPayments(user string, q paydom.PaymentPagedDTOQuery) (*cordom.PagedDTO[paydom.Payment], error) {
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

	var lst []paydom.Payment
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &lst); err != nil {
		return nil, err
	}

	dto := cordom.PagedDTO[paydom.Payment]{}
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

// CreatePayment
func (u *WalletDynamoDbRepository) CreatePayment(p *paydom.Payment) (*paydom.Payment, error) {
	dueDateTime, err := validator.DateParse(p.DueDate)
	if err != nil {
		return nil, err
	}

	p.Id = dueDateTime.Format("20060102") + validator.NewULID(time.Now())
	p.CreatedAt = time.Now().Format("2006-01-02T15:04:05-07:00")

	p.PK = "APP#" + u.app + "#USER#" + p.UserId
	p.SK = "APP#" + u.app + "#MOVT#" + p.Id

	p.GSI1PK = "APP#" + u.app + "#MOVT_STATUS#" + p.Status.String()
	p.GSI1SK = "APP#" + u.app + "#MOVT_DUEDATE#" + dueDateTime.Format("20060102")

	avs, err := attributevalue.MarshalMap(p)
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

//GetPaymentById
func (u *WalletDynamoDbRepository) GetPaymentById(user string, payment string) (*paydom.Payment, error) {
	out, err := u.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#USER#" + user,
			},
			"SK": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#MOVT#" + payment,
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
	var dto paydom.Payment
	if err := attributevalue.UnmarshalMap(out.Item, &dto); err != nil {
		return nil, err
	}
	return &dto, nil
}

// DeletePayment
func (u *WalletDynamoDbRepository) DeletePayment(user string, payment string) error {
	_, err := u.db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#USER#" + user,
			},
			"SK": &types.AttributeValueMemberS{
				Value: "APP#" + u.app + "#MOVT#" + payment,
			},
		},
		TableName: aws.String(u.table),
	})
	return err
}

// PutPayment
func (u *WalletDynamoDbRepository) PutPayment(p *paydom.Payment) (*paydom.Payment, error) {
	r, err := u.GetPaymentById(p.UserId, p.Id)
	if err != nil {
		return nil, err
	}

	if r == nil {
		return u.CreatePayment(p)
	}

	p.PK = "APP#" + u.app + "#USER#" + p.UserId
	p.SK = "APP#" + u.app + "#MOVT#" + p.Id

	p.GSI1PK = "APP#" + u.app + "#MOVT_STATUS#" + p.Status.String()
	p.GSI1SK = r.GSI1SK

	avs, err := attributevalue.MarshalMap(p)
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
func (u *WalletDynamoDbRepository) CreateTransactionDetail(i *paydom.TransactionDetail) (*paydom.TransactionDetail, error) {
	i.DetailId = validator.NewULID(time.Now())
	i.CreatedAt = time.Now().Format("2006-01-02T15:04:05-07:00")

	i.PK = "APP#" + u.app + "#MOVT#" + i.TransactionId
	i.SK = "APP#" + u.app + "#MOVT_DETAIL#" + i.DetailId

	avs, err := attributevalue.MarshalMap(i)
	if err != nil {
		return nil, err
	}

	if _, err := u.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(u.table),
		Item:      avs,
	}); err != nil {
		return nil, err
	}

	return i, nil
}

//ListTransactionDetails
func (u *WalletDynamoDbRepository) ListTransactionDetails(payment string, q cordom.PagedDTOQuery) (*cordom.PagedDTO[paydom.TransactionDetail], error) {
	cipher := fmt.Sprintf("%s%d", payment, q.Limit)
	k, err := dynamo.DecodePageToken(q.PageToken, cipher)
	if err != nil {
		return nil, err
	}

	var out *dynamodb.QueryOutput
	out, err = u.db.Query(context.TODO(), &dynamodb.QueryInput{
		KeyConditionExpression: aws.String("PK = :v and begins_with(SK, :x)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v": &types.AttributeValueMemberS{
				Value: "APP#wallet#MOVT#" + payment,
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

	var lst []paydom.TransactionDetail
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &lst); err != nil {
		return nil, err
	}

	dto := cordom.PagedDTO[paydom.TransactionDetail]{}
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

//GetTrasactionDetail
func (u *WalletDynamoDbRepository) GetTransactionDetail(transactionId string, detailId string) (*paydom.TransactionDetail, error) {
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
	var dto paydom.TransactionDetail
	if err := attributevalue.UnmarshalMap(out.Item, &dto); err != nil {
		return nil, err
	}
	return &dto, nil
}

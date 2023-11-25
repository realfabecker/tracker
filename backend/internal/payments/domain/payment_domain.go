package domain

import (
	"fmt"
	"time"

	cordom "github.com/realfabecker/wallet/internal/core/domain"
)

//  Payment model info
//	@Description	Payment information
type Payment struct {
	PK          string        `dynamodbav:"PK" json:"-"`
	SK          string        `dynamodbav:"SK" json:"-"`
	GSI1PK      string        `dynamodbav:"GSI1_PK" json:"-"`
	GSI1SK      string        `dynamodbav:"GSI1_SK" json:"-"`
	Id          string        `dynamodbav:"PaymentId" json:"id" validate:"required" example:"2023050701GXEH91YBVV40C1FK50S1P0KC"`
	UserId      string        `dynamodbav:"UserId" json:"userId" validate:"required" example:"realfabecker@gmail"`
	Description string        `dynamodbav:"Description" json:"description" validate:"required" example:"Supermercado"`
	Title       string        `dynamodbav:"Title" json:"title" validate:"required" example:"Supermercado"`
	Value       float64       `dynamodbav:"Value" json:"value" validate:"required,min=1" example:"200"`
	DueDate     string        `dynamodbav:"DueDate" json:"dueDate" validate:"required,ISO8601" example:"2023-05-07"`
	CreatedAt   string        `dynamodbav:"CreatedAt" json:"createdAt" example:"2023-04-07T16:45:30Z"`
	Status      PaymentStatus `dynamodbav:"Status" json:"status" validate:"oneof=paid cancelled pending" example:"paid"`
	Type        PaymentType   `dynamodbav:"Type" json:"type" validate:"oneof=expense income invoice detail" example:"expense"`
} //	@name	Payment

//  TransactionDetail model info
//	@Description	Invoice Detail information
type TransactionDetail struct {
	PK            string `dynamodbav:"PK" json:"-"`
	SK            string `dynamodbav:"SK" json:"-"`
	DetailId      string `dynamodbav:"DetailId" json:"id" example:"2023050701GXEH91YBVV40C1FK50S1P0KC"`
	TransactionId string `dynamodbav:"TransactionId" json:"transactionId" validate:"required" example:"2023050701GXEH91YBVV40C1FK50S1P0XD"`
	UserId        string `dynamodbav:"UserId" json:"userId" validate:"required" example:"realfabecker@gmail"`
	Description   string `dynamodbav:"Description" json:"description" validate:"required" example:"Supermercado"`
	Title         string `dynamodbav:"Title" json:"title" validate:"required" example:"Supermercado"`
	Value         uint16 `dynamodbav:"Value" json:"value" validate:"required,min=1" example:"200"`
	CreatedAt     string `dynamodbav:"CreatedAt" json:"createdAt" example:"2023-04-07T16:45:30Z"`
} //	@name	TransactionDetail

// PaymentPagedDTOQuery
type PaymentPagedDTOQuery struct {
	cordom.PagedDTOQuery
	Status  *PaymentStatus `query:"status" validate:"omitempty,oneof=paid cancelled pending"`
	DueDate *int32         `query:"due_date" validate:"omitempty"`
	Period  *PaymentPeriod `query:"period" validate:"omitempty,oneof=this_week this_month last_month next_month" example:"this_month"`
} //	@name	PaymentPagedDTOQuery

// GetDueDate
func (p PaymentPagedDTOQuery) GetDueDate() string {
	if p.Period != nil {
		return p.Period.Format()
	}
	if p.DueDate != nil {
		return fmt.Sprint(*p.DueDate)
	}
	return time.Now().Format("20060102")
}

// PaymentPeriod
type PaymentPeriod string

const (
	PaymentThisMonth    PaymentPeriod = "this_month"
	PaymentLastMonth    PaymentPeriod = "last_month"
	PaymentNextMonth    PaymentPeriod = "next_month"
	PaymentPeridUnknown PaymentPeriod = "unknown"
)

func (p PaymentPeriod) String() string {
	switch p {
	case PaymentThisMonth:
		return "this_month"
	case PaymentLastMonth:
		return "last_month"
	case PaymentNextMonth:
		return "next_month"
	}
	return "unknown"
}

// Format
func (p PaymentPeriod) Format() string {
	year, month, _ := time.Now().Date()
	switch p {
	case PaymentThisMonth:
		return time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Format("200601")
	case PaymentLastMonth:
		return time.Date(year, month-1, 1, 0, 0, 0, 0, time.Local).Format("200601")
	case PaymentNextMonth:
		return time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local).Format("200601")
	}
	return time.Now().Format("20060102")
}

// PaymentStatus
type PaymentStatus string //	@name	PaymentStatus

const (
	PaymentPaid      PaymentStatus = "paid"
	PaymentPending   PaymentStatus = "pending"
	PaymentCancelled PaymentStatus = "cancelled"
)

func (p PaymentStatus) String() string {
	switch p {
	case PaymentPaid:
		return "paid"
	case PaymentPending:
		return "pending"
	case PaymentCancelled:
		return "cancelled"
	}
	return "unknown"
}

// PaymentType
type PaymentType string //	@name	PaymentType

const (
	PaymentTypeExpense PaymentType = "expense"
	PaymentTypeIncome  PaymentType = "income"
)

func (p PaymentType) String() string {
	switch p {
	case PaymentTypeExpense:
		return "expense"
	case PaymentTypeIncome:
		return "income"
	}
	return "unknown"
}

package domain

import (
	"fmt"
	"time"
)

// Transaction model info
// @Description	Transaction information
type Transaction struct {
	TransactionId string            `json:"transactionId" validate:"required" example:"2023050701GXEH91YBVV40C1FK50S1P0KC"`
	UserId        string            `json:"userId" validate:"required" example:"e8ec3241-03b4-4aed-99d5-d72e1922d9b8"`
	Description   string            `json:"description" validate:"required" example:"Supermercado"`
	Title         string            `json:"title" validate:"required" example:"Supermercado"`
	Value         float64           `json:"value" validate:"required,min=1" example:"200"`
	DueDate       string            `json:"dueDate" validate:"required,ISO8601" example:"2023-05-07"`
	CreatedAt     string            `json:"createdAt" example:"2023-04-07T16:45:30Z"`
	Status        TransactionStatus `json:"status" validate:"oneof=paid cancelled pending" example:"paid"`
	Type          TransactionType   `json:"type" validate:"oneof=expense income invoice detail" example:"expense"`
} // @name	Transaction

//	TransactionDetail model info
//
// @Description	Invoice Detail information
type TransactionDetail struct {
	DetailId      string `json:"detailId" example:"2023050701GXEH91YBVV40C1FK50S1P0KC"`
	TransactionId string `json:"transactionId" validate:"required" example:"2023050701GXEH91YBVV40C1FK50S1P0XD"`
	UserId        string `json:"userId" validate:"required" example:"e8ec3241-03b4-4aed-99d5-d72e1922d9b8"`
	Description   string `json:"description" validate:"required" example:"Supermercado"`
	Title         string `json:"title" validate:"required" example:"Supermercado"`
	Value         uint16 `json:"value" validate:"required,min=1" example:"200"`
	CreatedAt     string `json:"createdAt" example:"2023-04-07T16:45:30Z"`
} //	@name	TransactionDetail

// TransactionPagedDTOQuery
type TransactionPagedDTOQuery struct {
	PagedDTOQuery
	DueDate *int32             `query:"due_date" validate:"omitempty" example:"2023"`
	Period  *TransactionPeriod `query:"period" validate:"omitempty,oneof=this_week this_month last_month next_month" example:"this_month"`
} //	@name	TransactionPagedDTOQuery

// GetDueDate
func (p TransactionPagedDTOQuery) GetDueDate() string {
	if p.Period != nil {
		return p.Period.Format()
	}
	if p.DueDate != nil {
		return fmt.Sprint(*p.DueDate)
	}
	return time.Now().Format("20060102")
}

// TransactionPeriod
type TransactionPeriod string

const (
	TransactionThisMonth    TransactionPeriod = "this_month"
	TransactionLastMonth    TransactionPeriod = "last_month"
	TransactionNextMonth    TransactionPeriod = "next_month"
	TransactionPeridUnknown TransactionPeriod = "unknown"
)

func (p TransactionPeriod) String() string {
	switch p {
	case TransactionThisMonth:
		return "this_month"
	case TransactionLastMonth:
		return "last_month"
	case TransactionNextMonth:
		return "next_month"
	}
	return "unknown"
}

// Format
func (p TransactionPeriod) Format() string {
	year, month, _ := time.Now().Date()
	switch p {
	case TransactionThisMonth:
		return time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Format("200601")
	case TransactionLastMonth:
		return time.Date(year, month-1, 1, 0, 0, 0, 0, time.Local).Format("200601")
	case TransactionNextMonth:
		return time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local).Format("200601")
	}
	return time.Now().Format("20060102")
}

// TransactionStatus
type TransactionStatus string // @name	TransactionStatus

const (
	TransactionPaid      TransactionStatus = "paid"
	TransactionPending   TransactionStatus = "pending"
	TransactionCancelled TransactionStatus = "cancelled"
)

func (p TransactionStatus) String() string {
	switch p {
	case TransactionPaid:
		return "paid"
	case TransactionPending:
		return "pending"
	case TransactionCancelled:
		return "cancelled"
	}
	return "unknown"
}

// TransactionType
type TransactionType string // @name TransactionType

const (
	TransactionTypeExpense TransactionType = "expense"
	TransactionTypeIncome  TransactionType = "income"
)

func (p TransactionType) String() string {
	switch p {
	case TransactionTypeExpense:
		return "expense"
	case TransactionTypeIncome:
		return "income"
	}
	return "unknown"
}

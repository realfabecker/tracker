package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/realfabecker/wallet/internal/adapters/common/validator"
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

type TransactionController struct {
	repository corpts.TransactionRepository
	service    corpts.TransactionService
	auth       corpts.AuthService
}

func NewTransactionController(
	walletRepository corpts.TransactionRepository,
	walletService corpts.TransactionService,
	auth corpts.AuthService,
) *TransactionController {
	return &TransactionController{walletRepository, walletService, auth}
}

// ListTransactions list transactions
//
//	@Summary		List transactions
//	@Description	List transactions
//	@Tags			Transactions
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			limit		query		number	true	"Number of records"
//	@Param			page_token	query		string	false	"Pagination token"
//	@Param			due_date	query		string	true	"Transaction due date"
//	@Success		200			{object}	cordom.ResponseDTO[cordom.PagedDTO[cordom.Transaction]]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/transactions [get]
func (w *TransactionController) ListTransactions(c *fiber.Ctx) error {
	q := cordom.TransactionPagedDTOQuery{}
	if err := c.QueryParser(&q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.NewValidator()
	if err := v.Struct(q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	out, err := w.service.ListTransactions(user.Subject, q)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(cordom.ResponseDTO[cordom.PagedDTO[cordom.Transaction]]{
		Status: "success",
		Data:   out,
	})
}

// GetTransactionById get a transaction by its id
//
//	@Summary		Get transaction by id
//	@Description	Get transaction by id
//	@Tags			Transactions
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			id	path		string	true	"Transaction id"
//	@Success		200	{object}	cordom.ResponseDTO[cordom.Transaction]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/transactions/{transactionId} [get]
func (w *TransactionController) GetTransactionById(c *fiber.Ctx) error {
	p := cordom.Transaction{}
	if err := c.ParamsParser(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.NewValidator()
	if err := v.StructPartial(p, "transactionId"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	d, err := w.service.GetTransactionById(user.Subject, p.TransactionId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if d == nil && err == nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	return c.JSON(cordom.ResponseDTO[cordom.Transaction]{
		Status: "success",
		Data:   d,
	})
}

// CreateTransaction create a new transaction
//
//	@Summary		Create a transaction
//	@Description	Create a new transaction record
//	@Tags			Transactions
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			request	body		cordom.Transaction	true	"Transaction payload"
//	@Success		200		{object}	cordom.ResponseDTO[cordom.Transaction]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/transactions [post]
func (w *TransactionController) CreateTransaction(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}
	body := cordom.Transaction{UserId: user.Subject}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	v := validator.NewValidator()
	if err := v.StructExcept(body, "TransactionId"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	p, err := w.service.CreateTransaction(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(cordom.ResponseDTO[cordom.Transaction]{
		Status: "success",
		Data:   p,
	})
}

// DeleteTransaction delete a transaction by its Id
//
//	@Summary		Delete transaction
//	@Description	Delete transaction
//	@Tags			Transactions
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			id	path		string	true	"Transaction id"
//	@Success		200	{object}	cordom.EmptyResponseDTO
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/transactions/{transactionId} [delete]
func (w *TransactionController) DeleteTransaction(c *fiber.Ctx) error {
	p := cordom.Transaction{}
	if err := c.ParamsParser(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.NewValidator()
	if err := v.StructPartial(p, "Id"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	transaction, err := w.service.GetTransactionById(user.Subject, p.TransactionId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if transaction == nil {
		return fiber.NewError(fiber.StatusNotFound, "Not Found")
	}

	if err := w.service.DeleteTransaction(user.Subject, p.TransactionId); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(204)
}

// PutTransaction
//
//	@Summary		Put a transaction
//	@Description	Update/Create a transaction record
//	@Tags			Transactions
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			id		path		string				true	"Transaction id"
//	@Param			request	body		cordom.Transaction	true	"Transaction payload"
//	@Success		200		{object}	cordom.ResponseDTO[cordom.Transaction]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/transactions/{transactionId} [put]
func (w *TransactionController) PutTransaction(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	body := cordom.Transaction{UserId: user.Subject}
	if err := c.ParamsParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	transaction, err := w.service.GetTransactionById(user.Subject, body.TransactionId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if transaction == nil {
		return fiber.NewError(fiber.StatusNotFound, "Not Found")
	}

	if err := c.BodyParser(&transaction); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.NewValidator()
	if err := v.Struct(transaction); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err = w.service.PutTransaction(transaction)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(204)
}

// CreateTransactionDetail
//
//	@Summary		Create a transaction detail
//	@Description	Create a new transaction detail record
//	@Tags			Transactions
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			request	body		cordom.TransactionDetail	true	"TransactionDetail payload"
//	@Success		200		{object}	cordom.ResponseDTO[cordom.TransactionDetail]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/transactions/{transactionId}/details [post]
func (w *TransactionController) CreateTransactionDetail(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	p := cordom.Transaction{}
	if err := c.ParamsParser(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	pv := validator.NewValidator()
	if err := pv.StructPartial(p, "Id"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := cordom.TransactionDetail{UserId: user.ID, TransactionId: p.TransactionId}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	bv := validator.NewValidator()
	if err := bv.Struct(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	d, err := w.service.CreateTransactionDetail(user.Subject, &body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(cordom.ResponseDTO[cordom.TransactionDetail]{
		Status: "success",
		Data:   d,
	})
}

// ListTransactionDetails
//
//	@Summary		List transaction details
//	@Description	List transaction details
//	@Tags			Transactions
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			limit		query		number	true	"Number of records"
//	@Param			page_token	query		string	false	"Pagination token"
//	@Success		200			{object}	cordom.ResponseDTO[cordom.PagedDTO[cordom.TransactionDetail]]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/transactions/{transactionId}/details [get]
func (w *TransactionController) ListTransactionDetails(c *fiber.Ctx) error {
	q := cordom.PagedDTOQuery{}
	if err := c.QueryParser(&q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.NewValidator()
	if err := v.Struct(q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	p := cordom.Transaction{}
	if err := c.ParamsParser(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	out, err := w.service.ListTransactionDetails(p.TransactionId, q)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(cordom.ResponseDTO[cordom.PagedDTO[cordom.TransactionDetail]]{
		Status: "success",
		Data:   out,
	})
}

// GetTransactionDetail
//
//	@Summary		Get transaction detail by id
//	@Description	Get transaction detail by id
//	@Tags			Transactions
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			transactionId	path		string	true	"Transaction id"
//	@Param			detailId		path		string	true	"Detail id"
//	@Success		200				{object}	cordom.ResponseDTO[cordom.TransactionDetail]
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/wallet/transactions/{transactionId}/details/{detailId} [get]
func (w *TransactionController) GetTransactionDetail(c *fiber.Ctx) error {
	p := cordom.TransactionDetail{}
	if err := c.ParamsParser(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.NewValidator()
	if err := v.StructPartial(p, "detailId", "transactionId"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	d, err := w.service.GetTransactionDetail(p.TransactionId, p.DetailId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if d == nil && err == nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	return c.JSON(cordom.ResponseDTO[cordom.TransactionDetail]{
		Status: "success",
		Data:   d,
	})
}

// DeleteTransactionDetail
//
//	@Summary		Delete a transaction detail by its id
//	@Description	Delete a transaction detail by its id
//	@Tags			Transactions
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			transactionId	path		string	true	"Transaction id"
//	@Param			detailId		path		string	true	"Detail id"
//	@Success		200				{object}	cordom.ResponseDTO[cordom.TransactionDetail]
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/wallet/transactions/{transactionId}/details/{detailId} [delete]
func (w *TransactionController) DeleteTransactionDetail(c *fiber.Ctx) error {
	p := cordom.TransactionDetail{}
	if err := c.ParamsParser(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.NewValidator()
	if err := v.StructPartial(p, "detailId", "transactionId"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	d, err := w.service.GetTransactionDetail(p.TransactionId, p.DetailId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if d == nil && err == nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	if err := w.service.DeleteTransactionDetail(p.TransactionId, p.DetailId); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to delete detail")
	}

	return c.JSON(cordom.ResponseDTO[cordom.TransactionDetail]{
		Status: "success",
		Data:   d,
	})
}

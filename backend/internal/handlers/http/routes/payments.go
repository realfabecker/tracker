package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/realfabecker/wallet/internal/adapters/common/validator"
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

// WalltetController definição de controlador http wallet
type WalletController struct {
	repository corpts.WalletRepository
	service    corpts.WalletService
	auth       corpts.AuthService
}

// NewWalletController construção de controlador http wallet
func NewWalletController(
	walletRepository corpts.WalletRepository,
	walletService corpts.WalletService,
	auth corpts.AuthService,
) *WalletController {
	return &WalletController{walletRepository, walletService, auth}
}

// ListUserPayments get user payments list
//
//	@Summary		List user payments
//	@Description	List user payments
//	@Tags			Payments
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			limit		query		number	true	"Number of records"
//	@Param			page_token	query		string	false	"Pagination token"
//	@Param			due_date	query		string	true	"Payment due date"
//	@Success		200			{object}	cordom.ResponseDTO[cordom.PagedDTO[cordom.Payment]]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/payments [get]
func (w *WalletController) ListUserPayments(c *fiber.Ctx) error {
	q := cordom.PaymentPagedDTOQuery{}
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

	out, err := w.service.ListPayments(user.Subject, q)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(cordom.ResponseDTO[cordom.PagedDTO[cordom.Payment]]{
		Status: "success",
		Data:   out,
	})
}

// ListUserPayments get user payments list
//
//	@Summary		Get payment by id
//	@Description	Get payment by id
//	@Tags			Payments
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			id	path		string	true	"Payment id"
//	@Success		200	{object}	cordom.ResponseDTO[cordom.Payment]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/payments/{id} [get]
func (w *WalletController) GetPaymentById(c *fiber.Ctx) error {
	p := cordom.Payment{}
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

	d, err := w.service.GetPaymentById(user.Subject, p.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if d == nil && err == nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	return c.JSON(cordom.ResponseDTO[cordom.Payment]{
		Status: "success",
		Data:   d,
	})
}

// CreateUserPayment get user information filtered by E-mail
//
//	@Summary		Create a payment
//	@Description	Create a new payment record
//	@Tags			Payments
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			request	body		cordom.Payment	true	"Payment payload"
//	@Success		200		{object}	cordom.ResponseDTO[cordom.Payment]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/payments [post]
func (w *WalletController) CreateUserPayment(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}
	body := cordom.Payment{UserId: user.Subject}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	v := validator.NewValidator()
	if err := v.StructExcept(body, "Id"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Not Found")
	}
	p, err := w.service.CreatePayment(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(cordom.ResponseDTO[cordom.Payment]{
		Status: "success",
		Data:   p,
	})
}

//	 DeletePaymentById
//
//		@Summary		Delete payment
//		@Description	Delete payment
//		@Tags			Payments
//		@Security		ApiKeyAuth
//		@Produce		json
//		@Param			id	path		string	true	"Payment id"
//		@Success		200	{object}	cordom.EmptyResponseDTO
//		@Failure		400
//		@Failure		500
//		@Router			/wallet/payments/{id} [delete]
func (w *WalletController) DeletePayment(c *fiber.Ctx) error {
	p := cordom.Payment{}
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

	payment, err := w.service.GetPaymentById(user.Subject, p.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if payment == nil {
		return fiber.NewError(fiber.StatusNotFound, "Not Found")
	}

	if err := w.service.DeletePayment(user.Subject, p.Id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(204)
}

// PutUserPayment
//
//	@Summary		Put a payment
//	@Description	Update/Create a payment record
//	@Tags			Payments
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			id		path		string			true	"Payment id"
//	@Param			request	body		cordom.Payment	true	"Payment payload"
//	@Success		200		{object}	cordom.ResponseDTO[cordom.Payment]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/payments [put]
func (w *WalletController) PutUserPayment(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	body := cordom.Payment{UserId: user.Subject}
	if err := c.ParamsParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	payment, err := w.service.GetPaymentById(user.Subject, body.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if payment == nil {
		return fiber.NewError(fiber.StatusNotFound, "Not Found")
	}

	if err := c.BodyParser(&payment); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.NewValidator()
	if err := v.Struct(payment); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err = w.service.PutPayment(payment)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(204)
}

// CreateTransactionDetail create transaction detail
//
//	@Summary		Create a transaction detail
//	@Description	Create a new transaction detail record
//	@Tags			Payments
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			request	body		cordom.TransactionDetail	true	"TransactionDetail payload"
//	@Success		200		{object}	cordom.ResponseDTO[cordom.TransactionDetail]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/payments/{id}/detail [post]
func (w *WalletController) CreateTransactionDetail(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	p := cordom.Payment{}
	if err := c.ParamsParser(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	pv := validator.NewValidator()
	if err := pv.StructPartial(p, "Id"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := cordom.TransactionDetail{UserId: user.ID, TransactionId: p.Id}
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

// ListTransasctionDetails list transaction details
//
//	@Summary		List transaction details
//	@Description	List transaction details
//	@Tags			Payments
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			limit		query		number	true	"Number of records"
//	@Param			page_token	query		string	false	"Pagination token"
//	@Success		200			{object}	cordom.ResponseDTO[cordom.PagedDTO[cordom.TransactionDetail]]
//	@Failure		400
//	@Failure		500
//	@Router			/wallet/payments/{id}/details [get]
func (w *WalletController) ListTransactionDetails(c *fiber.Ctx) error {
	q := cordom.PagedDTOQuery{}
	if err := c.QueryParser(&q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.NewValidator()
	if err := v.Struct(q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	p := cordom.Payment{}
	if err := c.ParamsParser(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	out, err := w.service.ListTransactionDetails(p.Id, q)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(cordom.ResponseDTO[cordom.PagedDTO[cordom.TransactionDetail]]{
		Status: "success",
		Data:   out,
	})
}

// Get transaction detail by id
//
//	@Summary		Get transaction detail by id
//	@Description	Get transaction detail by id
//	@Tags			Payments
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			transactionId	path		string	true	"Transaction id"
//	@Param			detailId	path		string	true	"Detail id"
//	@Success		200	{object}	cordom.ResponseDTO[cordom.TransactionDetail]
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/wallet/payments/{transactionId}/details/{detailId} [get]
func (w *WalletController) GetTransactionDetail(c *fiber.Ctx) error {
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

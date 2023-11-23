package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	"github.com/realfabecker/wallet/internal/core/lib/validator"
	usrdom "github.com/realfabecker/wallet/internal/users/domain"
	usrpts "github.com/realfabecker/wallet/internal/users/ports"
)

// WalltetController definição de controlador http wallet
type AuthController struct {
	authSrv usrpts.AuthService
	userSrv usrpts.UserService
}

// NewWalletController construção de controlador http wallet
func NewAuthController(
	authSrv usrpts.AuthService,
	userSrv usrpts.UserService,
) *AuthController {
	return &AuthController{authSrv, userSrv}
}

// UserLogin get user login by e-mal
//
//	@Summary		Get user login by e-mail
//	@Description	Get user login by e-mail
//	@Tags			Auth
//	@Produce		json
//	@Success		200	{object}	cordom.ResponseDTO[usrdom.UserToken]
//	@Failure		400
//	@Failure		500
//	@Router			/auth/login [get]
func (w *AuthController) Login(c *fiber.Ctx) error {
	q := usrdom.WalletLoginDTO{}
	if err := c.BodyParser(&q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.NewValidator()
	if err := v.Struct(q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	token, err := w.authSrv.Login(q.Email, q.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if token.AuthChallenge != nil {
		return c.Status(401).JSON(cordom.ResponseDTO[usrdom.UserToken]{
			Status: "error",
			Data:   token,
		})
	}

	return c.JSON(cordom.ResponseDTO[usrdom.UserToken]{
		Status: "success",
		Data:   token,
	})
}

// GetUserByEmail get user information filtered by E-mail
//
//	@Summary		Get user by e-mail
//	@Description	Get user information by e-mail
//	@Tags			Users
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Success		200	{object}	cordom.ResponseDTO[usrdom.User]
//	@Failure		400
//	@Failure		500
//	@Router			/users/profile [get]
func (w *AuthController) GetUserByEmail(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.RegisteredClaims)
	if !ok {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}
	out, err := w.userSrv.GetUserByEmail(user.Subject)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(cordom.ResponseDTO[usrdom.User]{
		Status: "success",
		Data:   out,
	})
}

//  Change User Password
//
//	@Summary		Change user password
//	@Description	Change user password
//	@Tags			Auth
//	@Produce		json
//	@Success		200	{object}	cordom.ResponseDTO[usrdom.UserToken]
//	@Failure		400
//	@Failure		500
//	@Router			/auth/change [get]
func (w *AuthController) Change(c *fiber.Ctx) error {
	q := usrdom.WalletLoginChangeDTO{}
	if err := c.BodyParser(&q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	v := validator.NewValidator()
	if err := v.Struct(q); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	token, err := w.authSrv.Change(q.Email, q.NewPassword, q.Session)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(cordom.ResponseDTO[usrdom.UserToken]{
		Status: "success",
		Data:   token,
	})
}

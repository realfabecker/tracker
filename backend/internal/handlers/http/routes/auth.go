package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/realfabecker/wallet/internal/adapters/common/validator"
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

type AuthController struct {
	authSrv corpts.AuthService
	userSrv corpts.UserService
}

func NewAuthController(
	authSrv corpts.AuthService,
	userSrv corpts.UserService,
) *AuthController {
	return &AuthController{authSrv, userSrv}
}

// Login user login
//
//	@Summary		User login
//	@Description	User login
//	@Tags			Auth
//	@Param			request	body	cordom.UserLoginDTO	true	"Login payload"
//	@Produce		json
//	@Success		200	{object}	cordom.ResponseDTO[cordom.UserToken]
//	@Failure		400
//	@Failure		500
//	@Router			/auth/login [post]
func (w *AuthController) Login(c *fiber.Ctx) error {
	q := cordom.UserLoginDTO{}
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
		return c.Status(401).JSON(cordom.ResponseDTO[cordom.UserToken]{
			Status: "error",
			Data:   token,
		})
	}

	return c.JSON(cordom.ResponseDTO[cordom.UserToken]{
		Status: "success",
		Data:   token,
	})
}

// Change Password
//
//	@Summary		Change password
//	@Description	Change password
//	@Tags			Auth
//	@Param			request	body	cordom.UserLoginChangeDTO	true	"Login payload"
//	@Produce		json
//	@Success		200	{object}	cordom.ResponseDTO[cordom.UserToken]
//	@Failure		400
//	@Failure		500
//	@Router			/auth/change [post]
func (w *AuthController) Change(c *fiber.Ctx) error {
	q := cordom.UserLoginChangeDTO{}
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

	return c.JSON(cordom.ResponseDTO[cordom.UserToken]{
		Status: "success",
		Data:   token,
	})
}

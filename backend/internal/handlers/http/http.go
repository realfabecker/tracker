package http

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/swagger"

	"github.com/realfabecker/wallet/internal/handlers/http/docs"

	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
	"github.com/realfabecker/wallet/internal/handlers/http/routes"
)

type Handler struct {
	app              *fiber.App
	walletConfig     *cordom.Config
	walletController *routes.TransactionController
	authService      corpts.AuthService
}

//	@title						Wallet Rest API
//	@version					1.0
//	@description				Wallet Rest API
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Type 'Bearer ' and than your API token
func NewFiberHandler(
	walletConfig *cordom.Config,
	walletController *routes.TransactionController,
	authService corpts.AuthService,
) corpts.HttpHandler {

	// open api base project configuration (2)
	docs.SwaggerInfo.Host = walletConfig.AppHost
	docs.SwaggerInfo.Schemes = []string{"http"}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			msgs := utils.StatusMessage(code)

			var ferr *fiber.Error
			if errors.As(err, &ferr) {
				code = ferr.Code
				msgs = ferr.Message
			}

			c.Status(code)
			return c.JSON(cordom.ResponseDTO[interface{}]{
				Status:  "error",
				Message: msgs,
				Code:    code,
			})
		},
	})
	return &Handler{
		app,
		walletConfig,
		walletController,
		authService,
	}
}

func (a *Handler) GetApp() interface{} {
	return a.app
}

func (a *Handler) Listen(port string) error {
	return a.app.Listen(":" + port)
}

func (a *Handler) Register() error {
	a.app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 30 * time.Second,
	}))

	a.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	a.app.Get("/docs/*", swagger.HandlerDefault)
	wallet := a.app.Group("/wallet")

	tran := wallet.Group("/transactions")
	tran.Use(a.authHandler)
	tran.Post("/", a.walletController.CreateTransaction)
	tran.Get("/", a.walletController.ListTransactions)
	tran.Get("/:transactionId", a.walletController.GetTransactionById)
	tran.Delete("/:transactionId", a.walletController.DeleteTransaction)
	tran.Put("/:transactionId", a.walletController.PutTransaction)
	tran.Post("/:transactionId/details", a.walletController.CreateTransactionDetail)
	tran.Get("/:transactionId/details", a.walletController.ListTransactionDetails)
	tran.Get("/:transactionId/details/:detailId", a.walletController.GetTransactionDetail)
	return nil
}

func (a *Handler) authHandler(c *fiber.Ctx) error {
	auth := c.Get("authorization")
	if len(auth) < (len("bearer") + 1) {
		return fiber.NewError(fiber.ErrUnauthorized.Code)
	}
	u, err := a.authService.Verify(auth[len("bearer")+1:])
	if err != nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, err.Error())
	}
	c.Locals("user", u)
	return c.Next()
}

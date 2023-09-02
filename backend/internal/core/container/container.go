package container

import (
	"context"

	awsconf "github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/realfabecker/wallet/internal/core/lib/env"

	"github.com/realfabecker/wallet/internal/payments/ports"
	payrep "github.com/realfabecker/wallet/internal/payments/repositories/dynamodb"
	paysrv "github.com/realfabecker/wallet/internal/payments/services"

	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
	corsrv "github.com/realfabecker/wallet/internal/core/services"
	inthld "github.com/realfabecker/wallet/internal/handlers/http"
	payhld "github.com/realfabecker/wallet/internal/payments/handler/http"
	usrhld "github.com/realfabecker/wallet/internal/users/handler/http"
	usrpts "github.com/realfabecker/wallet/internal/users/ports"
	usrrep "github.com/realfabecker/wallet/internal/users/repositories/dynamodb"
	usrsrv "github.com/realfabecker/wallet/internal/users/services"

	"go.uber.org/dig"
)

var Container dig.Container

// init
func init() {
	Container = *dig.New()

	Container.Provide(func() (*cordom.Config, error) {
		cnf := &cordom.Config{}
		if err := env.Unmarshal(cnf); err != nil {
			return nil, err
		}
		return cnf, nil
	})

	Container.Provide(func(cnf *cordom.Config) (*dynamodb.Client, error) {
		env, err := awsconf.LoadDefaultConfig(context.TODO())
		if err != nil {
			return nil, err
		}
		return dynamodb.NewFromConfig(env), nil
	})

	Container.Provide(func() corpts.CacheHandler {
		return corsrv.NewFileCache()
	})

	Container.Provide(func(cache corpts.CacheHandler) corpts.JwtHandler {
		return corsrv.NewJwtHandler(cache)
	})

	Container.Provide(func(cnf *cordom.Config) (*cognito.Client, error) {
		env, err := awsconf.LoadDefaultConfig(context.TODO())
		if err != nil {
			return nil, err
		}

		return cognito.NewFromConfig(env), nil
	})

	Container.Provide(func(d *dynamodb.Client, cnf *cordom.Config) (ports.WalletRepository, error) {
		return payrep.NewWalletDynamoDBRepository(d, cnf.DynamoDBTableName, cnf.AppName)
	})

	Container.Provide(func(r ports.WalletRepository) ports.WalletService {
		return paysrv.NewWalletService(r)
	})

	Container.Provide(func(
		walletConfig *cordom.Config,
		cognitoClient *cognito.Client,
		jwtHandler corpts.JwtHandler,
	) usrpts.AuthService {
		return usrsrv.NewCognitoAuthService(
			walletConfig.CognitoClientId,
			walletConfig.CognitoJwkUrl,
			cognitoClient,
			jwtHandler,
		)
	})

	Container.Provide(func(
		r ports.WalletRepository,
		s ports.WalletService,
		t usrpts.AuthService,
	) (*payhld.WalletController, error) {
		return payhld.NewWalletController(r, s, t), nil
	})

	Container.Provide(func(d *dynamodb.Client, cnf *cordom.Config) (usrpts.UserRepository, error) {
		return usrrep.NewUserRepository(d, cnf.DynamoDBTableName, cnf.AppName)
	})

	Container.Provide(func(
		r usrpts.UserRepository,
	) (usrpts.UserService, error) {
		return usrsrv.NewUserService(r), nil
	})

	Container.Provide(func(
		a usrpts.AuthService,
		u usrpts.UserService,
	) (*usrhld.AuthController, error) {
		return usrhld.NewAuthController(a, u), nil
	})

	Container.Provide(func(
		walletConfig *cordom.Config,
		walletController *payhld.WalletController,
		usersController *usrhld.AuthController,
		authService usrpts.AuthService,
	) (corpts.HttpHandler, error) {
		return inthld.NewFiberHandler(
			walletConfig,
			walletController,
			usersController,
			authService,
		), nil
	})
}

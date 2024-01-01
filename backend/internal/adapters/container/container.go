package container

import (
	"context"
	"github.com/gofiber/fiber/v2/log"

	awsconf "github.com/aws/aws-sdk-go-v2/config"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/realfabecker/wallet/internal/adapters/common/cache"
	"github.com/realfabecker/wallet/internal/adapters/common/dotenv"
	"github.com/realfabecker/wallet/internal/adapters/common/jwt2"

	payrep "github.com/realfabecker/wallet/internal/adapters/transactions/repositories"
	usrsrv "github.com/realfabecker/wallet/internal/adapters/users/services"
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
	corsrv "github.com/realfabecker/wallet/internal/core/services"
	"github.com/realfabecker/wallet/internal/handlers/http"
	"github.com/realfabecker/wallet/internal/handlers/http/routes"
	"go.uber.org/dig"
)

var Container dig.Container

func init() {
	Container = *dig.New()
	if err := reg3(); err != nil {
		log.Fatalf("unable to register services %v", err)
	}
}

func reg3() error {

	if err := Container.Provide(func() (*cordom.Config, error) {
		cnf := &cordom.Config{}
		if err := dotenv.Unmarshal(cnf); err != nil {
			return nil, err
		}
		return cnf, nil
	}); err != nil {
		return err
	}

	if err := Container.Provide(func(cnf *cordom.Config) (*dynamodb.Client, error) {
		env, err := awsconf.LoadDefaultConfig(context.TODO())
		if err != nil {
			return nil, err
		}
		return dynamodb.NewFromConfig(env), nil
	}); err != nil {
		return err
	}

	if err := Container.Provide(func() corpts.CacheHandler {
		return cache.NewFileCache()
	}); err != nil {
		return err
	}

	if err := Container.Provide(func() corpts.JwtHandler {
		return jwt2.NewJwtHandler()
	}); err != nil {
		return err
	}

	if err := Container.Provide(func(cnf *cordom.Config) (*cognito.Client, error) {
		env, err := awsconf.LoadDefaultConfig(context.TODO())
		if err != nil {
			return nil, err
		}
		return cognito.NewFromConfig(env), nil
	}); err != nil {
		return err
	}

	if err := Container.Provide(func(d *dynamodb.Client, cnf *cordom.Config) (corpts.TransactionRepository, error) {
		return payrep.NewWalletDynamoDBRepository(d, cnf.DynamoDBTableName, cnf.AppName)
	}); err != nil {
		return err
	}

	if err := Container.Provide(func(r corpts.TransactionRepository) corpts.TransactionService {
		return corsrv.NewTransactionService(r)
	}); err != nil {
		return err
	}

	if err := Container.Provide(func(
		jwtHandler corpts.JwtHandler,
	) corpts.AuthService {
		return usrsrv.NewCognitoAuthService(jwtHandler)
	}); err != nil {
		return err
	}

	if err := Container.Provide(func(
		r corpts.TransactionRepository,
		s corpts.TransactionService,
		t corpts.AuthService,
	) (*routes.TransactionController, error) {
		return routes.NewTransactionController(r, s, t), nil
	}); err != nil {
		return err
	}

	if err := Container.Provide(func(
		walletConfig *cordom.Config,
		walletController *routes.TransactionController,
		authService corpts.AuthService,
	) (corpts.HttpHandler, error) {
		return http.NewFiberHandler(
			walletConfig,
			walletController,
			authService,
		), nil
	}); err != nil {
		return err
	}

	return nil
}

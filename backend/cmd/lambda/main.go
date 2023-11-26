package main

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/gofiber/fiber/v2"
	"github.com/realfabecker/wallet/internal/adapters/container"

	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

// Lambda callback handler
func handler(flb *fiberadapter.FiberLambda) func(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	return func(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
		response, err := flb.ProxyWithContext(ctx, events.APIGatewayProxyRequest{
			Path:                  request.RequestContext.HTTP.Path,
			Body:                  request.Body,
			Headers:               request.Headers,
			HTTPMethod:            request.RequestContext.HTTP.Method,
			QueryStringParameters: request.QueryStringParameters,
			IsBase64Encoded:       request.IsBase64Encoded,
		})

		if err != nil {
			return events.LambdaFunctionURLResponse{}, err
		}

		if response.Headers == nil {
			response.Headers = map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "*",
				"Access-Control-Allow-Headers": "*",
			}
		} else {
			response.Headers["Access-Control-Allow-Origin"] = "*"
			response.Headers["Access-Control-Allow-Methods"] = "*"
			response.Headers["Access-Control-Allow-Headers"] = "*"
		}

		return events.LambdaFunctionURLResponse{
			StatusCode: response.StatusCode,
			Headers:    response.Headers,
			Body:       response.Body,
		}, nil
	}
}

func main() {
	if err := container.Container.Invoke(func(
		app corpts.HttpHandler,
		walletConfig *cordom.Config,
	) error {
		if err := app.Register(); err != nil {
			return err
		}

		fapp, b := app.GetApp().(*fiber.App)
		if !b {
			return errors.New("not possible to capture fiber instance in lambda")
		}

		lambda.Start(handler(fiberadapter.New(fapp)))

		return nil
	}); err != nil {
		log.Fatalln(err)
	}
}

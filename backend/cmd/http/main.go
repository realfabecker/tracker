package main

import (
	"log"

	"github.com/realfabecker/wallet/internal/core/container"
	cordom "github.com/realfabecker/wallet/internal/core/domain"
	corpts "github.com/realfabecker/wallet/internal/core/ports"
)

func main() {
	if err := container.Container.Invoke(func(
		app corpts.HttpHandler,
		walletConfig *cordom.Config,
	) error {
		if err := app.Register(); err != nil {
			return err
		}
		return app.Listen(walletConfig.AppPort)
	}); err != nil {
		log.Fatalln(err)
	}
}

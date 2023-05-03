package main

import (
	"user-services/infrastructure/app"
	"user-services/utils"

	"github.com/rs/zerolog"
)

func main() {
	// utility setup
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := utils.Logger()

	// application instance setup
	apps := app.AppInstance(logger)
	apps.LoadEnv()
	apps.Run()
}

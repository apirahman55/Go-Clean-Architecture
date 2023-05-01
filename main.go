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

	// why i'm define the spesific of database name?
	// yea, i think there's a people need more database instance
	// and you can setup that instance more than once here
	dbMysql := apps.MysqlConn()
}

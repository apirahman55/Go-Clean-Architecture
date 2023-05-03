package app

import (
	"user-services/domain/entities"
	"user-services/infrastructure/app/routers"
	"user-services/infrastructure/database"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func AppInstance(logger *zerolog.Logger) *appInstance {
	return &appInstance{logger}
}

type appInstance struct {
	logger *zerolog.Logger
}

func (app *appInstance) LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		defer panic(err)
		app.logger.WithLevel(zerolog.PanicLevel).Msg(err.Error())

		return
	}

	app.logger.Info().Msg("Env Loaded Succesfully...")
}

func (app *appInstance) mysqlConn() *gorm.DB {
	db, err := database.ConnMysqlDb()
	if err != nil {
		defer panic(err)
		app.logger.WithLevel(zerolog.PanicLevel).Msg(err.Error())

		return nil
	}

	app.logger.Info().Msg("Database Connected...")
	db.AutoMigrate(&entities.User{})
	return db
}

func (app *appInstance) Run() {
	e := echo.New()
	db := app.mysqlConn()

	router := routers.NewAppRoute(e, db, app.logger)
	router.UserRouter()

	e.Logger.Fatal(e.Start(":3333"))
}

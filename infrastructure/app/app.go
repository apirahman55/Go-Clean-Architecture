package app

import (
	"user-services/infrastructure/database"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func AppInstance(logger *zerolog.Logger) *App {
	return &App{logger}
}

type App struct {
	logger *zerolog.Logger
}

func (app *App) LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		defer panic(err)
		app.logger.WithLevel(zerolog.PanicLevel).Msg(err.Error())

		return
	}

	app.logger.Info().Msg("Env Loaded Succesfully...")
}

func (app *App) MysqlConn() *gorm.DB {
	db, err := database.ConnMysqlDb()
	if err != nil {
		defer panic(err)
		app.logger.WithLevel(zerolog.PanicLevel).Msg(err.Error())

		return nil
	}

	app.logger.Info().Msg(err.Error())
	return db
}

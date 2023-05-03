package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func NewAppRoute(echo *echo.Echo, db *gorm.DB, logger *zerolog.Logger) *appRoute {
	return &appRoute{echo, db, logger}
}

type appRoute struct {
	echo   *echo.Echo
	db     *gorm.DB
	logger *zerolog.Logger
}

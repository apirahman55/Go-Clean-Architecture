package service

import "github.com/labstack/echo/v4"

type UserService interface {
	GetAllUser(c echo.Context) error
}

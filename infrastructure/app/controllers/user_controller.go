package controllers

import (
	"user-services/domain/service"

	"github.com/labstack/echo/v4"
)

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *userController {
	return &userController{service}
}

func (c *userController) GetAllUser(context echo.Context) error {
	return c.service.GetAllUser(context)
}

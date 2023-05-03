package services

import (
	"net/http"
	"user-services/domain/dtos"
	"user-services/domain/repository"
	"user-services/domain/service"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func NewUserService(repo repository.UserRepository, logger *zerolog.Logger) service.UserService {
	return &userService{repo, logger}
}

type userService struct {
	repo   repository.UserRepository
	logger *zerolog.Logger
}

func (s *userService) GetAllUser(c echo.Context) error {
	users, err := s.repo.FindAll()
	if err != nil {
		res := dtos.MetadataDto{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, res)
	}

	res := dtos.GetUsersDto{
		Data: users,
		Meta: dtos.MetadataDto{
			Status:  http.StatusOK,
			Message: "Success",
		},
	}

	return c.JSON(http.StatusOK, res)
}

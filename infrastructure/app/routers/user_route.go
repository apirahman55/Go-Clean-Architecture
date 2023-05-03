package routers

import (
	"user-services/infrastructure/app/controllers"
	"user-services/infrastructure/app/services"
	"user-services/infrastructure/repositories"
)

func (r *appRoute) UserRouter() {
	repository := repositories.NewUserRepository(r.db)
	controller := controllers.NewUserController(services.NewUserService(repository, r.logger))

	r.echo.GET("/users", controller.GetAllUser)
}

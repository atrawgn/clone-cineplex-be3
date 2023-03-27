package route

import (
	"github.com/atrawiguna/golang-restapi-gorm/config"
	"github.com/atrawiguna/golang-restapi-gorm/controller"
	"github.com/atrawiguna/golang-restapi-gorm/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("/login", controller.LoginController)

	r.Get("/film", middleware.Auth, controller.FilmControllerGet)
	r.Get("/film/:id", controller.FilmControllerGetById)
	r.Post("/film", controller.FilmControllerCreate)
	r.Put("/film/:id", controller.FilmControllerUpdate)
	r.Delete("/film/:id", controller.FilmControllerDelete)

	r.Get("/user", middleware.Auth, controller.UserControllerGet)
	r.Get("/user/:id", controller.UserControllerGetById)
	r.Post("/user", controller.UserControllerCreate)
	r.Put("/user/:id", controller.UserControllerUpdate)
	r.Delete("/user/:id", controller.UserControllerDelete)
}

package routes

import (
	"pg-conn/src/controllers"

	"github.com/gofiber/fiber"
)

func ConfigRoutes(r *fiber.App) *fiber.App {
	router := r.Group("api/v1")
	{
		router.Get("/", controllers.ListAllUsers)
		router.Post("/create", controllers.CreateUser)
		router.Put("/update/id::id", controllers.UpdateUser)
		router.Patch("/updatepass/id::id", controllers.UpdateUserPassword)
		router.Delete("/delete/id::id", controllers.Delete)
	}

	return r
}

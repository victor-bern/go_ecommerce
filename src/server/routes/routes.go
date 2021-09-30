package routes

import (
	"pg-conn/src/controllers"

	"github.com/gofiber/fiber"
)

func ConfigRoutes(r *fiber.App) *fiber.App {
	router := r.Group("api/v1")
	{
		router.Group("user")
		{
			router.Get("/", controllers.ListAllUsers)
			router.Post("/create", controllers.CreateUser)
			router.Put("/update/id::id", controllers.UpdateUser)
			router.Patch("/updatepass/id::id", controllers.UpdateUserPassword)
			router.Delete("/delete/id::id", controllers.Delete)
		}

		router.Group("product")
		{
			router.Get("/all", controllers.GetAllproduct)
			router.Get("/getbyid/id::id")
			router.Get("/getbytitle/title::title", controllers.GetProductByTitle)
			router.Post("/create", controllers.InsertProduct)
			router.Put("/update/id::id")
			router.Get("/delete/id::id")
		}

	}

	return r
}

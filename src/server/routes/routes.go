package routes

import (
	"pg-conn/src/controllers"

	"github.com/gofiber/fiber"
)

func ConfigRoutes(r *fiber.App) *fiber.App {
	router := r.Group("api/v1")
	{
		user := router.Group("/user")
		{
			user.Get("/", controllers.ListAllUsers)
			user.Post("/create", controllers.CreateUser)
			user.Put("/update/id::id", controllers.UpdateUser)
			user.Patch("/updatepass/id::id", controllers.UpdateUserPassword)
			user.Delete("/delete/id::id", controllers.Delete)
		}

		product := router.Group("/product")
		{
			product.Get("/all", controllers.GetAllproduct)
			product.Get("/getbyid/id::id", controllers.GetProductById)
			product.Get("/getbytitle/title::title", controllers.GetProductByTitle)
			product.Post("/create", controllers.InsertProduct)
			product.Put("/update/id::id", controllers.UpdateProduct)
			//product.Get("/delete/id::id")
		}

	}

	return r
}

package controllers

import (
	"pg-conn/src/database"
	"pg-conn/src/models"
	"pg-conn/src/repositories"

	"github.com/gofiber/fiber"
)

func GetAllproduct(ctx *fiber.Ctx) {
	productRepo := repositories.NewProductRepo(database.GetDatabase())

	products, err := productRepo.All()
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	ctx.Status(200).JSON(products)
}

func GetProductByTitle(ctx *fiber.Ctx) {
	var product models.Product
	productRepo := repositories.NewProductRepo(database.GetDatabase())
	title := ctx.Params("title")

	err := productRepo.GetByTitle(title, &product)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	ctx.Status(200).JSON(product)
}

func GetProductById(ctx *fiber.Ctx) {
	var product models.Product
	productRepo := repositories.NewProductRepo(database.GetDatabase())
	id := ctx.Params("id")

	err := productRepo.GetById(id, &product)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	ctx.Status(200).JSON(product)
}

func InsertProduct(ctx *fiber.Ctx) {
	var product models.ProductRequest
	productRepo := repositories.NewProductRepo(database.GetDatabase())

	err := ctx.BodyParser(&product)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	newProduct := models.Product{
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
	}

	_, err = productRepo.InserProduct(newProduct, product.Inventory)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	ctx.Status(200).JSON(fiber.Map{
		"Message": "product created with success",
	})
}

func UpdateProduct(ctx *fiber.Ctx) {
	var product models.Product
	pr := repositories.NewProductRepo(database.GetDatabase())
	id := ctx.Params("id")

	err := ctx.BodyParser(&product)

	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	err = pr.UpdateProduct(product, id)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	ctx.Status(200).JSON(fiber.Map{
		"Message": "product updated with success",
	})
}

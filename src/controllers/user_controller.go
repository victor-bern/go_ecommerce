package controllers

import (
	"pg-conn/src/database"
	"pg-conn/src/models"
	"pg-conn/src/repositories"

	"github.com/gofiber/fiber"
)

func ListAllUsers(ctx *fiber.Ctx) {
	db := database.GetDatabase()

	userRepo := repositories.NewUserRepo(db)

	users, err := userRepo.All()

	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	ctx.Status(fiber.StatusOK).JSON(users)
}

func CreateUser(ctx *fiber.Ctx) {
	db := database.GetDatabase()
	userRepo := repositories.NewUserRepo(db)
	var user models.User

	err := ctx.BodyParser(&user)
	if err != nil {
		ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	userFinded := userRepo.GetByEmail(user.Email)
	if err != nil {
		ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	if len(userFinded.Email) > 0 {
		ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": "This user email already exists",
		})
		return
	}

	_, err = userRepo.Create(user)

	if err != nil {
		ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "User created with success",
	})
}

func UpdateUser(ctx *fiber.Ctx) {
	db := database.GetDatabase()
	id := ctx.Params("id")
	var user models.User
	userRepo := repositories.NewUserRepo(db)

	err := ctx.BodyParser(&user)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	_, err = userRepo.Update(user, id)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "User updated",
	})
}

func UpdateUserPassword(ctx *fiber.Ctx) {
	db := database.GetDatabase()
	id := ctx.Params("id")
	var user models.User
	userRepo := repositories.NewUserRepo(db)

	err := ctx.BodyParser(&user)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}

	err = userRepo.UpdatePassword(user.Password, id)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err.Error(),
		})
		return
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "User password updated",
	})
}

func Delete(ctx *fiber.Ctx) {
	db := database.GetDatabase()
	id := ctx.Params("id")
	userRepo := repositories.NewUserRepo(db)
	err := userRepo.DeleteUser(id)
	if err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	ctx.Status(200)
}

package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/josepedro1903/go-rest-api/database"
	"github.com/josepedro1903/go-rest-api/models"
)

func Home(ctx *fiber.Ctx) error {
	fmt.Fprint(ctx, "Home page")

	return nil
}

func TodasPersonalidades(ctx *fiber.Ctx) error {
	var p []models.Personalidade
	database.DB.Find(&p)

	return ctx.JSON(p)
}

func RetornaPersonalidade(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	return ctx.JSON(personalidade)
}

func CriaUmaNovaPersonalidade(ctx *fiber.Ctx) error {
	var novaPersonalidade models.Personalidade

	ctx.BodyParser(&novaPersonalidade)

	database.DB.Create(&novaPersonalidade)

	return ctx.JSON(novaPersonalidade)
}

func DeletaUmaPersonalidade(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var personalidade models.Personalidade

	database.DB.Delete(&personalidade, id)

	return ctx.JSON(personalidade)
}

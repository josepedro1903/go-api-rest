package controllers

import (
	"encoding/json"
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

	json.NewEncoder(ctx).Encode(p)

	return nil
}

func RetornaPersonalidade(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	json.NewEncoder(ctx).Encode(personalidade)

	return nil
}

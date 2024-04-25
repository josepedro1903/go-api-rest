package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/josepedro1903/go-rest-api/models"
)

func Home(ctx *fiber.Ctx) error {
	fmt.Fprint(ctx, "Home page")

	return nil
}

func TodasPersonalidades(ctx *fiber.Ctx) error {
	json.NewEncoder(ctx).Encode(models.Personalidades)

	return nil
}

func RetornaPersonalidade(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	IdInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("ID inválido")
	}

	for _, personalidade := range models.Personalidades {
		if personalidade.ID == IdInt {
			return ctx.JSON(personalidade)
		}
	}

	return ctx.Status(fiber.StatusNotFound).SendString("Personalidade não encontrada")
}

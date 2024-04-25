package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/josepedro1903/go-rest-api/models"
	"github.com/josepedro1903/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{ID: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Starting server...")

	app := fiber.New()
	routes.HandleRequest(app)
	app.Listen(":8080")
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josepedro1903/go-rest-api/controllers"
)

func HandleRequest(app fiber.Router) {
	app.Get("/", controllers.Home)
	app.Get("/personalidades", controllers.TodasPersonalidades)
	app.Get("/personalidades/:id", controllers.RetornaPersonalidade)
}

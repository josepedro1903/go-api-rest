package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/josepedro1903/go-rest-api/controllers"
)

func HandleRequest(app fiber.Router) {

	r := app.Group("/api")

	r.Get("/", controllers.Home)
	r.Get("/personalidades", controllers.TodasPersonalidades)
	r.Get("/personalidades/:id", controllers.RetornaPersonalidade)

	r.Post("/personalidades", controllers.CriaUmaNovaPersonalidade)

}

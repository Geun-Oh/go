package router

import (
	"github.com/Geun-Oh/fiber-react/server/controller"
	"github.com/gofiber/fiber/v2"
)

// Setup Routes info
func SetUpRoutes(app *fiber.App) {

	// list => get
	// add => post
	// update => put
	// delete => delete

	app.Get("/", controller.BlogList)
	app.Post("/", controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/:id", controller.BlogDelete)
}

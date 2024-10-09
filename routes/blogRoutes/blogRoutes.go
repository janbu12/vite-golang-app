package blogroutes

import (
	"github.com/gofiber/fiber/v2"
	blogmodel "github.com/janbu12/vite-golang-app/models/blogModel"
)

func BlogRoutes(app *fiber.App) {
	app.Get("/blogs", blogmodel.GetAll)
	app.Get("/blogs/:id", blogmodel.GetOne)
	app.Post("/blogs", blogmodel.Create)
}

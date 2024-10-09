package main

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	blogroutes "github.com/janbu12/vite-golang-app/routes/blogRoutes"
	"github.com/janbu12/vite-golang-app/sup"
	"github.com/janbu12/vite-golang-app/ui"
)

func main() {
	app := fiber.New()

	sup.InitSupabase()

	index, err := fs.Sub(ui.Index, "dist")
	if err != nil {
		log.Fatal(err)
	}

	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(index),
		Index:  "index.html",
		Browse: false,
	}))

	serverUI := func(c *fiber.Ctx) error {
		return filesystem.SendFile(c, http.FS(index), "index.html")
	}

	app.Get("/login", serverUI)

	blogroutes.BlogRoutes(app)

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	log.Fatal(app.Listen(":8000"))
}

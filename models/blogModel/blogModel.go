package blogmodel

import (
	"github.com/gofiber/fiber/v2"
	"github.com/janbu12/vite-golang-app/entities"
	"github.com/janbu12/vite-golang-app/sup"
)

func GetAll(c *fiber.Ctx) error {
	var blogs []entities.Blog
	err := sup.Client.DB.From("blog").Select("*").Execute(&blogs)
	if err != nil {
		panic(err)
	}
	return c.Status(200).JSON(blogs)
}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var blog []entities.Blog
	err := sup.Client.DB.From("blog").Select("*").Eq("id", id).Execute(&blog)
	if err != nil {
		panic(err)
	}
	return c.Status(200).JSON(blog)
}

func Create(c *fiber.Ctx) error {
	var blog entities.FormBlog
	err := c.BodyParser(&blog)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	var results []entities.Blog
	err = sup.Client.DB.From("blog").Insert(blog).Execute(&results)
	if err != nil {
		panic(err)
	}
	return c.Status(201).JSON(blog)
}

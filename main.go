package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./templates", ".html")
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)
	app.Static("/public", "./public")
	app.Get("/", mainPage)
	app.Listen(":8000")
}

func mainPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"name": c.Query("name", "World"),
	})
}

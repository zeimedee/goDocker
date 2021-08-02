package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World")
	})

	app.Listen(":4000")
}

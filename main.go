package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name") + "!")
		}
		return c.SendString("Where is John?")
	})

	app.Listen(":3000")

}

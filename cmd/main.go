package main

import (
	"log"
	"nateloeffel-dev-backend/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	app := fiber.New()

	// Init session/middleware store
	// Sessions store in memory by default.
	// Set CookieSecure to true for HTTPS in prod.
	store := session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   false,
	})

	// stores sessions for each request
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("sessionStore", store)
		return c.Next()
	})

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

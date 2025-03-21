package routes

import (
	"nateloeffel-dev-backend/internal/handlers"
	"nateloeffel-dev-backend/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Example route.
	app.Get("/", handlers.HomeHandler)

	// Auth routes
	authGroup := app.Group("/auth")
	authGroup.Post("/login", handlers.LoginHandler)
	authGroup.Get("/logout", handlers.LogoutHandler)

	userGroup := app.Group("/user")
	userGroup.Use(middlewares.RequireAuth) // ensure session is valid
	userGroup.Get("/profile", handlers.ProfileHandler)

}

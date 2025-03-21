package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func RequireAuth(c *fiber.Ctx) error {
	store := c.Locals("sessionStore").(*session.Store)
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	isLoggedIn := sess.Get("isLoggedIn")
	if isLoggedIn == nil || isLoggedIn == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized - Please log in",
		})
	}

	return c.Next()
}

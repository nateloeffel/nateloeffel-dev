package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func ProfileHandler(c *fiber.Ctx) error {
	store := c.Locals("sessionStore").(*session.Store)
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	// retreive protected data
	username := sess.Get("username")
	return c.JSON(fiber.Map{
		"username":   username,
		"profileMsg": "This is your protected profile data.",
	})
}

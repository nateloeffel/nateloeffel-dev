package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *fiber.Ctx) error {

	// retrieve session store from locals
	store := c.Locals("sessionStore").(*session.Store)

	// Parse JSON
	var loginReq LoginRequest
	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// simple user/pass validation
	if loginReq.Username != "admin" || loginReq.Password != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Get/create session
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	// set session variables
	sess.Set("username", loginReq.Username)
	sess.Set("isLoggedIn", true)

	// Save session
	if err := sess.Save(); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Logged in successfully",
	})
}

func LogoutHandler(c *fiber.Ctx) error {
	store := c.Locals("sessionStore").(*session.Store)
	sess, err := store.Get(c)

	if err != nil {
		return err
	}

	// Delete Session
	if err := sess.Destroy(); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}

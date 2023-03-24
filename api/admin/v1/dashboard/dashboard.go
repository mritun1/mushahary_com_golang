package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mritun1/mushahary_com_golang.git/controllers/auth"
)

func Dash(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status": "You are now logged in, to protected page",
		"Name":   auth.GetLoggedInDetails(c, "name"),
		"Id":     auth.GetLoggedInDetails(c, "id"),
		"status": fiber.StatusOK,
	})
}

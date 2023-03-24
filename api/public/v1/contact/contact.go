package contact

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mritun1/mushahary_com_golang.git/controllers/e_mail"
)

func EmailSending(c *fiber.Ctx) error {
	FormData := new(ModalMail)
	c.BodyParser(FormData)
	status, er := e_mail.SendMail(FormData.EMAIL, "mritunjoy@72dragons.com", FormData.NAME+"--"+FormData.SUB+"--"+FormData.EMAIL, FormData.MESSAGE)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  er,
		"status": status,
	})
}

type ModalMail struct {
	NAME    string
	EMAIL   string
	SUB     string
	MESSAGE string
}

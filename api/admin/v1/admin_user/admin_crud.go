package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mritun1/mushahary_com_golang.git/controllers/auth"
	"github.com/mritun1/mushahary_com_golang.git/database/db"
	"github.com/mritun1/mushahary_com_golang.git/database/models"
)

func Create(c *fiber.Ctx) error {
	usr := new(models.User)
	if err := c.BodyParser(usr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	if usr.AVATAR == "" || usr.USER_NAME == "" || usr.PASSWORD == "" || usr.FIRST_NAME == "" || usr.LAST_NAME == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Please don't leave any field empty",
			"status": fiber.StatusBadRequest,
		})
	}

	//CHECK IF THE USER ALREADY EXISTS
	var userModal models.User
	db.Con.Where(&models.User{USER_NAME: usr.USER_NAME}).First(&userModal)

	if userModal.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "User Already Exists",
			"status": fiber.StatusBadRequest,
		})
	}

	//CONVERTING PASSWORD INTO HASH
	hash, _ := auth.HashPassword(usr.PASSWORD)
	usr.PASSWORD = hash

	db.Con.Create(&usr)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   "User Registered Success",
		"status":  fiber.StatusOK,
		"Details": usr,
	})
}

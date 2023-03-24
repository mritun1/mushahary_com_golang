package dash_photos

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mritun1/mushahary_com_golang.git/controllers/auth"
	"github.com/mritun1/mushahary_com_golang.git/database/db"
	"github.com/mritun1/mushahary_com_golang.git/database/models"
)

func GetPhotosByUser(c *fiber.Ctx) error {
	offset := c.Params("offset")
	limit := c.Params("limit")
	var photos []models.JoinPhotosList
	var count int64
	result := db.Con.Model(&models.Photos{}).Select("ROW_NUMBER() OVER () as sl, photos.*, photo_categories.category_name").Joins("left join photo_categories on photo_categories.id = photos.photo_category_id WHERE photos.user_id =" + auth.GetLoggedInDetails(c, "id") + "  LIMIT " + offset + "," + limit).Scan(&photos).Count(&count)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Sorry, No result found",
			"status": fiber.StatusBadRequest,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Found result",
		"status": fiber.StatusOK,
		"total":  count,
		"data":   photos,
	})
}

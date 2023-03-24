package dash_articles

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mritun1/mushahary_com_golang.git/controllers/auth"
	"github.com/mritun1/mushahary_com_golang.git/database/db"
	"github.com/mritun1/mushahary_com_golang.git/database/models"
)

func GetArticlesByUser(c *fiber.Ctx) error {
	offset := c.Params("offset")
	limit := c.Params("limit")
	var articles []models.JoinArticleList
	var count int64
	result := db.Con.Model(&models.Articles{}).Select("ROW_NUMBER() OVER () as sl, articles.*, categories.category_name").Joins("left join categories on categories.id = articles.category_id WHERE articles.user_id =" + auth.GetLoggedInDetails(c, "id") + " order by articles.id desc LIMIT " + offset + "," + limit).Scan(&articles).Count(&count)
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
		"data":   articles,
	})
}

package articles

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mritun1/mushahary_com_golang.git/database/db"
	"github.com/mritun1/mushahary_com_golang.git/database/models"
)

func GetById(c *fiber.Ctx) error {

	id := c.Params("id")

	var articles models.JoinArticleList

	result := db.Con.Model(&models.Articles{}).Select(" articles.*, categories.category_name").Joins("left join categories on categories.id = articles.category_id WHERE articles.id=" + id).Scan(&articles)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Sorry, No result found",
			"status": fiber.StatusBadRequest,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Found result",
		"status": fiber.StatusOK,
		"data":   articles,
	})
}
func GetAll(c *fiber.Ctx) error {
	offset := c.Params("offset")
	limit := c.Params("limit")
	var articles []models.JoinArticleList
	var articlesAll []models.Articles
	resultAll := db.Con.Find(&articlesAll)

	result := db.Con.Model(&models.Articles{}).Select("ROW_NUMBER() OVER () as sl, articles.*, categories.category_name").Joins("left join categories on categories.id = articles.category_id order by articles.id desc LIMIT " + offset + "," + limit).Scan(&articles)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Sorry, No result found",
			"status": fiber.StatusBadRequest,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Found result",
		"status": fiber.StatusOK,
		"total":  resultAll.RowsAffected,
		"data":   articles,
	})
}
func GetRelatedArticles(c *fiber.Ctx) error {
	limit := c.Params("limit")
	cat := c.Params("cat")
	var articles []models.JoinArticleList
	var articlesAll []models.Articles
	resultAll := db.Con.Find(&articlesAll)

	result := db.Con.Model(&models.Articles{}).Select("ROW_NUMBER() OVER () as sl, articles.*, categories.category_name").Joins("left join categories on categories.id = articles.category_id WHERE articles.category_id=" + cat + " order by articles.id desc  LIMIT " + limit).Scan(&articles)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Sorry, No result found",
			"status": fiber.StatusBadRequest,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Found result",
		"status": fiber.StatusOK,
		"total":  resultAll.RowsAffected,
		"data":   articles,
	})
}
func GetArticlesByCat(c *fiber.Ctx) error {
	limit := c.Params("limit")
	offset := c.Params("offset")
	cat := c.Params("cat")
	var articles []models.JoinArticleList
	var articlesAll []models.Articles
	resultAll := db.Con.Find(&articlesAll)

	result := db.Con.Model(&models.Articles{}).Select("ROW_NUMBER() OVER () as sl, articles.*, categories.category_name").Joins("left join categories on categories.id = articles.category_id WHERE articles.category_id=" + cat + " order by articles.id desc  LIMIT " + offset + "," + limit).Scan(&articles)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Sorry, No result found",
			"status": fiber.StatusBadRequest,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Found result",
		"status": fiber.StatusOK,
		"total":  resultAll.RowsAffected,
		"data":   articles,
	})
}
func SearchArticles(c *fiber.Ctx) error {
	//limit := c.Params("limit")
	//offset := c.Params("offset")
	search := c.Params("search")
	s := strings.ReplaceAll(search, "%20", " ")
	var articles []models.JoinArticleList
	var articlesAll []models.Articles
	resultAll := db.Con.Find(&articlesAll)

	result := db.Con.Model(&models.Articles{}).Select("ROW_NUMBER() OVER () as sl, articles.*, categories.category_name").Where("articles.title LIKE ? or articles.des LIKE ?", "%"+s+"%", "%"+s+"%").Joins("left join categories on categories.id = articles.category_id ").Limit(21).Scan(&articles)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Sorry, No result found",
			"status": fiber.StatusBadRequest,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Found result",
		"status": fiber.StatusOK,
		"total":  resultAll.RowsAffected,
		"data":   articles,
	})
}
func GetArticlesCat(c *fiber.Ctx) error {
	var cat []models.Category
	result := db.Con.Find(&cat)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Result OK",
		"status": fiber.StatusOK,
		"total":  result.RowsAffected,
		"data":   cat,
	})
}

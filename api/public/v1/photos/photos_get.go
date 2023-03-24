package photos

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mritun1/mushahary_com_golang.git/database/db"
	"github.com/mritun1/mushahary_com_golang.git/database/models"
)

func GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var photo models.JoinPhotosList

	result := db.Con.Model(&models.Photos{}).Select("photos.*, photo_categories.category_name").Joins("left join photo_categories on photo_categories.id = photos.photo_category_id  WHERE photos.id=" + id).Scan(&photo)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Sorry, No result found",
			"status": fiber.StatusBadRequest,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Found result",
		"status": fiber.StatusOK,
		"data":   photo,
	})
}
func GetAll(c *fiber.Ctx) error {
	offset := c.Params("offset")
	limit := c.Params("limit")
	var photos []models.JoinPhotosList
	//result := db.Con.Order("id DESC").Find(&photos)
	var count int64
	result := db.Con.Model(&models.Photos{}).Select("ROW_NUMBER() OVER () as sl, photos.*, photo_categories.category_name").Joins("left join photo_categories on photo_categories.id = photos.photo_category_id  LIMIT " + offset + "," + limit).Scan(&photos).Count(&count)
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
func GetRelatedPhotos(c *fiber.Ctx) error {
	limit := c.Params("limit")
	cat := c.Params("cat")
	var photos []models.JoinPhotosList
	var count int64
	result := db.Con.Model(&models.Photos{}).Select("ROW_NUMBER() OVER () as sl, photos.*, photo_categories.category_name").Joins("left join photo_categories on photo_categories.id = photos.photo_category_id WHERE photos.photo_category_id=" + cat + "  LIMIT " + limit).Scan(&photos).Count(&count)
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
func SearchPhotos(c *fiber.Ctx) error {
	//offset := c.Params("offset")
	//limit := c.Params("limit") + offset + "," + limit
	search := c.Params("search")
	s := strings.ReplaceAll(search, "%20", " ")
	var photos []models.JoinPhotosList
	//result := db.Con.Order("id DESC").Find(&photos)
	var count int64
	result := db.Con.Model(&models.Photos{}).Select("ROW_NUMBER() OVER () as sl, photos.*, photo_categories.category_name").Where("photos.photo_title LIKE ? or photos.photo_des LIKE ?", "%"+s+"%", "%"+s+"%").Joins("left join photo_categories on photo_categories.id = photos.photo_category_id  ").Limit(21).Scan(&photos).Count(&count)
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
func GetPhotosByCat(c *fiber.Ctx) error {
	limit := c.Params("limit")
	offset := c.Params("offset")
	cat := c.Params("cat")
	var photos []models.JoinPhotosList
	var count int64
	result := db.Con.Model(&models.Photos{}).Select("ROW_NUMBER() OVER () as sl, photos.*, photo_categories.category_name").Joins("left join photo_categories on photo_categories.id = photos.photo_category_id WHERE photos.photo_category_id=" + cat + "  LIMIT " + offset + "," + limit).Scan(&photos).Count(&count)
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
func GetPhotosCat(c *fiber.Ctx) error {
	var cat []models.Photo_category
	result := db.Con.Find(&cat)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Result OK",
		"status": fiber.StatusOK,
		"total":  result.RowsAffected,
		"data":   cat,
	})
}

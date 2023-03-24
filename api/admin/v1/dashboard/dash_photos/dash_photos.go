package dash_photos

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mritun1/mushahary_com_golang.git/controllers/auth"
	"github.com/mritun1/mushahary_com_golang.git/controllers/aws_s3"
	"github.com/mritun1/mushahary_com_golang.git/database/db"
	"github.com/mritun1/mushahary_com_golang.git/database/models"
)

func Create(c *fiber.Ctx) error {

	photos := new(models.Photos)
	//GETTING THE PHOTOS PARSER BODY
	if err := c.BodyParser(photos); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}

	//FILE UPLOAD TO S3
	result, err := aws_s3.Uploading(c, "testing1290", "mushahary_com/photos/original/")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	photos.PHOTO_FILE = result.Location

	//GET LOGGED IN USER_ID
	photos.USER_ID = auth.GetLoggedInDetails(c, "id")
	//INSERT TO DATABASE
	db.Con.Create(&photos)
	//RETURN THE OUTPUT
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Photos Added Success",
		"status": fiber.StatusOK,
		"id":     photos.ID,
	})
}
func Edit(c *fiber.Ctx) error {
	photos := new(models.Photos)
	//GETTING THE PHOTOS PARSER BODY
	if err := c.BodyParser(photos); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":  "Sorry! Bad request.",
			"status": fiber.StatusBadRequest,
		})
	}
	//UPDATE TO DATABASE
	db.Con.Model(&models.Photos{}).Where("user_id", auth.GetLoggedInDetails(c, "id")).Updates(&photos)
	//RETURN THE OUTPUT
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Photos Updated Success",
		"status": fiber.StatusOK,
		"id":     photos.ID,
	})
}
func Delete(c *fiber.Ctx) error {
	//CONVERTING PARAMETER INTO INT
	id, _ := strconv.Atoi(c.Params("id"))
	//CREATING A MODAL
	var photo models.Photos
	//COVETING INT INTO UINT AND SETTING MODAL ID
	photo.ID = uint(id)
	//CHECK IF DELETED SUCCESS OR NOT
	err := db.Con.First(&photo).Error
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":  "Sorry! Photo not Found",
			"status": fiber.StatusNotFound,
		})
	}
	//DELETING PHOTO FROM S3
	file_loc := strings.Replace(photo.PHOTO_FILE, "https://testing1290.s3.ap-south-1.amazonaws.com/", "", 1)
	err = aws_s3.DeleteS3("testing1290", file_loc)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	//DELETE QUERY
	db.Con.Where("user_id", auth.GetLoggedInDetails(c, "id")).Delete(&photo)
	//RETURN THE OUTPUT
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Deleted Success",
		"status": fiber.StatusOK,
	})
}

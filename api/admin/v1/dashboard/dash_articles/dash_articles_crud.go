package dash_articles

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mritun1/mushahary_com_golang.git/controllers/auth"
	"github.com/mritun1/mushahary_com_golang.git/controllers/aws_s3"
	"github.com/mritun1/mushahary_com_golang.git/database/db"
	"github.com/mritun1/mushahary_com_golang.git/database/models"
)

// --------------------------------------------------------------------------------------------------------
//
//	API FOR ARTICLE CREATE - START
//
// --------------------------------------------------------------------------------------------------------
func Create(c *fiber.Ctx) error {
	article := new(models.Articles)
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":  "Sorry! Bad request",
			"status": fiber.StatusBadRequest,
		})
	}

	//FILE UPLOAD TO S3
	result, err := aws_s3.Uploading(c, "testing1290", "mushahary_com/articles/thumbnail/original/")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	article.THUMBNAIL = result.Location

	//TARGETING USER ID
	article.USER_ID = auth.GetLoggedInDetails(c, "id")
	//INSERTING TO DATABASE
	db.Con.Create(&article)
	//RETURN WHEN INSERTED SUCCESS
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Articles Created Success",
		"status": fiber.StatusOK,
		"id":     article.ID,
	})
}

// --------------------------------------------------------------------------------------------------------
//
//	API FOR ARTICLE CREATE - END
//
// --------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------------------
//
//	API FOR ARTICLE UPDATE - START
//
// --------------------------------------------------------------------------------------------------------
func Edit(c *fiber.Ctx) error {
	article := new(models.Articles)

	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":  "Sorry! Bad request",
			"status": fiber.StatusBadRequest,
		})
	}
	//UPDATE TO DATABASE
	db.Con.Model(&models.Articles{}).Where("user_id=? and id=?", auth.GetLoggedInDetails(c, "id"), article.ID).Limit(1).Updates(&article)
	//RETURN WHEN INSERTED SUCCESS
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Articles UPDATED Success",
		"status": fiber.StatusOK,
		"id":     article.ID,
	})
}

// --------------------------------------------------------------------------------------------------------
//
//	API FOR ARTICLE UPDATE - END
//
// --------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------------------
//
//	API FOR ARTICLE DELETE - START
//
// --------------------------------------------------------------------------------------------------------
func Delete(c *fiber.Ctx) error {
	//CONVERTING PARAMETER INTO INT
	id, _ := strconv.Atoi(c.Params("id"))
	//CREATING A MODAL
	var article models.Articles
	//COVETING INT INTO UINT AND SETTING MODAL ID
	article.ID = uint(id)
	//CHECK IF DELETED SUCCESS OR NOT
	err := db.Con.First(&article).Error
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":  "Sorry! Article not Found",
			"status": fiber.StatusNotFound,
		})
	}
	//DELETING PHOTO FROM S3
	file_loc := strings.Replace(article.THUMBNAIL, "https://testing1290.s3.ap-south-1.amazonaws.com/", "", 1)
	err = aws_s3.DeleteS3("testing1290", file_loc)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":  err.Error(),
			"status": fiber.StatusBadRequest,
		})
	}
	//DELETE QUERY
	db.Con.Where("user_id", auth.GetLoggedInDetails(c, "id")).Delete(&article)
	//RETURN THE OUTPUT
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  "Articles Deleted Success",
		"status": fiber.StatusOK,
	})
}

//--------------------------------------------------------------------------------------------------------
//	API FOR ARTICLE DELETE - END
//--------------------------------------------------------------------------------------------------------

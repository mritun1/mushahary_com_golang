package admin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/mritun1/mushahary_com_golang.git/controllers/auth"
	"github.com/mritun1/mushahary_com_golang.git/database/db"
	"github.com/mritun1/mushahary_com_golang.git/database/models"
)

func Login(c *fiber.Ctx) error {

	FormData := new(models.User)
	c.BodyParser(FormData)

	var userMOdal models.User
	db.Con.Where(&models.User{USER_NAME: FormData.USER_NAME}).First(&userMOdal)

	if userMOdal.ID == 0 {
		//USER DOESN'T EXISTS
		return c.JSON(fiber.Map{
			"error":    fiber.StatusUnauthorized,
			"status":   "Sorry! Your Password or User Name doesn't exists",
			"response": 0,
		})
	}

	//USER IS EXISTS

	//CHECKING IF THE PASSWORD MATCH
	match := auth.CheckPasswordHash(FormData.PASSWORD, userMOdal.PASSWORD)
	if !match {
		return c.JSON(fiber.Map{
			"error":    fiber.StatusUnauthorized,
			"status":   "Sorry! Wrong Password",
			"response": 0,
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  userMOdal.FIRST_NAME + " " + userMOdal.LAST_NAME,
		"id":    fmt.Sprintf("%v", userMOdal.ID),
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(), //1 week expiry time
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(auth.SecretWord))
	if err != nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{"Status": fiber.StatusInternalServerError})
	}
	return c.JSON(fiber.Map{
		"error":    fiber.StatusOK,
		"status":   "Logged in success",
		"response": 1,
		"token":    t,
		"details": fiber.Map{
			"first_name": userMOdal.FIRST_NAME,
			"last_name":  userMOdal.LAST_NAME,
			"avatar":     userMOdal.AVATAR,
		},
	})

}

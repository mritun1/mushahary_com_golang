package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	admin "github.com/mritun1/mushahary_com_golang.git/api/admin/v1/admin_user"
	"github.com/mritun1/mushahary_com_golang.git/api/admin/v1/dashboard"
	"github.com/mritun1/mushahary_com_golang.git/api/admin/v1/dashboard/dash_articles"
	"github.com/mritun1/mushahary_com_golang.git/api/admin/v1/dashboard/dash_photos"
	"github.com/mritun1/mushahary_com_golang.git/api/public/v1/articles"
	"github.com/mritun1/mushahary_com_golang.git/api/public/v1/contact"
	"github.com/mritun1/mushahary_com_golang.git/api/public/v1/photos"
	"github.com/mritun1/mushahary_com_golang.git/controllers/auth"
	"github.com/mritun1/mushahary_com_golang.git/database/db"
)

func main() {
	//app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024, // this is the Uploading limit of 20MB
	})

	// Default config
	//app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8001",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// HOME
	app.Get("/", func(c *fiber.Ctx) error {
		var word string
		word = "Database Connected success"
		if db.Err != nil {
			word = db.Err.Error()
		}
		return c.SendString(word)
	})

	// PUBLIC ARTICLES
	app.Get("/api/v1/articles/getAll/:offset/:limit", articles.GetAll)
	app.Get("/api/v1/articles/getById/:id", articles.GetById)
	app.Get("/api/v1/articles/cat", articles.GetArticlesCat)
	app.Get("/api/v1/articles/getRelatedArticles/:limit/:cat", articles.GetRelatedArticles)
	app.Get("/api/v1/articles/getArticlesByCat/:offset/:limit/:cat", articles.GetArticlesByCat)
	app.Get("/api/v1/articles/searchArticles/:offset/:limit/:search", articles.SearchArticles)

	// PUBLIC PHOTOS
	app.Get("/api/v1/photos/getAll/:offset/:limit", photos.GetAll)
	app.Get("/api/v1/photos/searchPhotos/:offset/:limit/:search", photos.SearchPhotos)
	app.Get("/api/v1/photos/getById/:id", photos.GetById)
	app.Get("/api/v1/photos/cat", photos.GetPhotosCat)
	app.Get("/api/v1/photos/relatedPhotos/:cat/:limit", photos.GetRelatedPhotos)
	app.Get("/api/v1/photos/getPhotosByCat/:offset/:limit/:cat", photos.GetPhotosByCat)

	// ADMIN USER
	app.Post("/api/v1/admin/create", admin.Create)
	app.Post("/api/v1/admin/login", admin.Login)

	//CONTACT US
	app.Post("/email", contact.EmailSending)

	//--------------------------------------------------------
	//--------------------------------------------------------

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			authHeader := c.Get("Authorization")
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"error":  "Unauthorized",
				"status": fiber.StatusUnauthorized,
				"key":    authHeader,
			})
		},
		SigningKey: []byte(auth.SecretWord),
	}))

	//--------------------------------------------------------
	//--------------------------------------------------------

	// ADMIN DASH
	app.Get("/api/v1/admin/dash", dashboard.Dash)

	// ADMIN ARTICLES CRUD
	app.Post("/api/v1/articles/create", dash_articles.Create)
	app.Put("/api/v1/articles/edit", dash_articles.Edit)
	app.Delete("/api/v1/articles/delete/:id", dash_articles.Delete)
	app.Get("/api/v1/articles/only_admin/:offset/:limit", dash_articles.GetArticlesByUser)

	// ADMIN PHOTOS CRUD
	app.Post("/api/v1/photos/create", dash_photos.Create)
	app.Put("/api/v1/photos/edit", dash_photos.Edit)
	app.Delete("/api/v1/photos/delete/:id", dash_photos.Delete)
	app.Get("/api/v1/photos/only_admin/:offset/:limit", dash_photos.GetPhotosByUser)

	app.Listen(os.Getenv("LOCAL_PORT"))
}

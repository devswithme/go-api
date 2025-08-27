package app

import (
	c "fyque/controller/auth"
	m "fyque/middleware/auth"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouter(app *fiber.App, db *gorm.DB, authController c.AuthController){
	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)

	auth.Use(m.NewAuthMiddleware(db))
}
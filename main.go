package main

import (
	"fyque/app"
	authController "fyque/controller/auth"
	authRepo "fyque/repository/auth"
	authService "fyque/service/auth"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main(){
	db, err := app.InitDB()

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	
	authRepo := authRepo.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepo)
	authController := authController.NewAuthController(authService)
	
	web := fiber.New(fiber.Config{
		AppName: "fyque",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.NewRouter(web, db, authController)

	web.Listen(":3000")
}
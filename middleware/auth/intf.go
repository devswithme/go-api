package auth

import "github.com/gofiber/fiber/v2"

type AuthMiddleware interface {
	Handler() fiber.Handler
}
package auth

import (
	"fmt"
	"fyque/helper"
	"fyque/model/domain"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthMiddlewareImpl struct {
	db *gorm.DB

}

func NewAuthMiddleware(db *gorm.DB) AuthMiddleware {
	return &AuthMiddlewareImpl{
		db: db,
	}
}

func (m *AuthMiddlewareImpl) Handler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr, err := helper.ExtractToken(c)

		if err != nil {
			c.ClearCookie("jwt")

			// !TODO
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error){
			if t.Method.Alg() != jwt.GetSigningMethod("HS256").Alg() {
				return nil, fmt.Errorf("%d", t.Header["Alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.ClearCookie("jwt")

			// !TODO
		}

		userId := token.Claims.(jwt.MapClaims)["userId"]

		user := new(domain.User)

		if err := m.db.Where("id = ?", userId).First(&user).Error; err != nil {
			c.ClearCookie("jwt")

			// !TODO
		}

		c.Locals("userId", userId)
		return c.Next()
	}
}
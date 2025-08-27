package helper

import (
	"errors"
	"fyque/model/domain"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *domain.User) (string, error){
	secret := []byte("super-secret-key")
		method := jwt.SigningMethodHS256
		claims := jwt.MapClaims{
			"userId": user.ID,
			"username": user.Username,
			"exp": time.Now().Add(time.Hour * 168).Unix(),
		}

		token, err := jwt.NewWithClaims(method, claims).SignedString(secret)

		if err != nil {
			return "", err
		}

		return token, nil
}

func ExtractToken(c *fiber.Ctx) (string, error){
	if cookie := c.Cookies("jwt"); cookie != "" {
		return cookie, nil
	}

	if authHeader := c.Get("autorization"); authHeader != "" {
		return strings.TrimPrefix(authHeader, "Bearer "), nil
	}

	return "", errors.New("missing token")
	
}
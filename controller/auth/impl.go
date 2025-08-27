package auth

import (
	"fyque/model/domain"
	"fyque/model/web"
	"fyque/service/auth"
	"log"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	authService auth.AuthService
}

func NewAuthController(authService auth.AuthService) AuthController {
	return &AuthControllerImpl{
		authService: authService,
	}
}

func (ctr *AuthControllerImpl) Register(c *fiber.Ctx) error {
	user := new(domain.User)

	
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	token, err := ctr.authService.Register(web.RegisterRequest{
		Username: user.Username,
		Password: user.Password,
	})

	log.Default().Println(err)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (ctr *AuthControllerImpl) Login(c *fiber.Ctx) error{
	user := new(domain.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	token, err := ctr.authService.Login(web.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

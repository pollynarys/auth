package handlers

import (
	"auth/model"
	"auth/repository"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, err := encrypt(data)
	if err != nil {
		return err
	}

	user := model.User{
		Name:              data["name"],
		Email:             data["email"],
		EncryptedPassword: password,
	}

	repository.DB.Create(user)

	return c.JSON(user)
}

func encrypt(data map[string]string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	var user model.User

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	repository.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if !MatchPassword(user.EncryptedPassword, data["password"]) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	return c.JSON(user)
}

func MatchPassword(correctPassword string, enteredPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(correctPassword), []byte(enteredPassword)) == nil
}

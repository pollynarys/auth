package handlers

import (
	"auth/model"
	"auth/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

var SecretKey = "secret"

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

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func MatchPassword(correctPassword string, enteredPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(correctPassword), []byte(enteredPassword)) == nil
}

func GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	var user model.User

	claims := token.Claims.(*jwt.StandardClaims) // convert interface Claims with only method valid in struct StandardClaims with Issuer
	repository.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

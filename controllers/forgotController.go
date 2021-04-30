package controllers

import (
	"math/rand"
	"net/smtp"

	"github.com/backend/database"
	"github.com/backend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Forgot password method
func Forgot(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	token := RandomStringRunes(12)

	passwordReset := models.PasswordReset{
		Email: data["email"],
		Token: token,
	}

	database.DB.Create(&passwordReset)

	from := "matias@example.com"

	to := []string{
		data["email"],
	}

	// Vue is 8080 port by defualt.
	url := "http://localhost:8080/reset/" + token

	message := []byte("Click <a href=\"" + url + "\">here</a> to reset your password")

	// This is mailhog. https://github.com/mailhog/MailHog - just easy and fast way to test if my mail system works.
	err := smtp.SendMail("0.0.0:1025", nil, from, to, message)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

// Reset password method
func Reset(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}

	var passwordReset = models.PasswordReset{}

	if err := database.DB.Where("token = ?", data["token"]).Last(&passwordReset); err.Error != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	database.DB.Model(&models.User{}).Where("email = ?", passwordReset.Email).Update("password", password)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

// RandomStringRunes : Creates a random string that we can use as token
func RandomStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

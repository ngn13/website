package util

import (
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func TitleToID(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", ""))
}

func CreateToken() string {
	s := make([]byte, 32)
	for i := 0; i < 32; i++ {
		s[i] = byte(65 + rand.Intn(25))
	}
	return string(s)
}

func ErrorCheck(err error, c *fiber.Ctx) bool {
	if err != nil {
		log.Printf("Server error: '%s' on %s\n", err, c.Path())
		return true
	}

	return false
}

func ErrorJSON(error string) fiber.Map {
	return fiber.Map{
		"error": error,
	}
}

func GetIP(c *fiber.Ctx) string {
	if c.Get("X-Real-IP") != "" {
		return strings.Clone(c.Get("X-Real-IP"))
	}

	return c.IP()
}

func ErrServer(c *fiber.Ctx) error {
	return c.Status(http.StatusInternalServerError).JSON(ErrorJSON("Server error"))
}

func ErrExists(c *fiber.Ctx) error {
	return c.Status(http.StatusConflict).JSON(ErrorJSON("Entry already exists"))
}

func ErrBadData(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(ErrorJSON("Provided data is invalid"))
}

func ErrBadJSON(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(ErrorJSON("Bad JSON data"))
}

func ErrAuth(c *fiber.Ctx) error {
	return c.Status(http.StatusUnauthorized).JSON(ErrorJSON("Authentication failed"))
}

func NoError(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(ErrorJSON(""))
}

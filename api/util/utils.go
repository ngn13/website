package util

import (
	"crypto/sha512"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetSHA512(s string) string {
	hasher := sha512.New()
	return fmt.Sprintf("%x", hasher.Sum([]byte(s)))
}

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

func ErrEntryExists(c *fiber.Ctx) error {
	return c.Status(http.StatusConflict).JSON(ErrorJSON("Entry already exists"))
}

func ErrEntryNotExists(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(ErrorJSON("Entry does not exist"))
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

func ErrNotFound(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(ErrorJSON("Requested endpoint not found"))
}

func NoError(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(ErrorJSON(""))
}

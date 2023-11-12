package util

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var charlist  []rune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func CreateToken() string {
  b := make([]rune, 20)
  for i := range b {
    b[i] = charlist[rand.Intn(len(charlist))]
  }

  return string(b)
}

func ErrorCheck(err error, c *fiber.Ctx) bool{
  if err != nil {
    log.Printf("Server error: '%s' on %s\n", err, c.Path())
    return true 
  }

  return false 
}

func ErrorJSON(error string) fiber.Map{
  return fiber.Map {
    "error": error,
  }
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


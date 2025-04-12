package util

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/config"
)

func IP(c *fiber.Ctx) string {
	conf := c.Locals("config").(*config.Type)

	if conf.IPHeader != "" && c.Get(conf.IPHeader) != "" {
		return strings.Clone(c.Get(conf.IPHeader))
	}

	return c.IP()
}

func JSON(c *fiber.Ctx, code int, data fiber.Map) error {
	if data == nil {
		data = fiber.Map{}
		data["error"] = ""
	} else if _, ok := data["error"]; !ok {
		data["error"] = ""
	}

	if data["error"] == 200 {
		Warn("200 response with an error at %s", c.Path())
	}

	return c.Status(code).JSON(data)
}

func ErrInternal(c *fiber.Ctx, err error) error {
	Warn("Internal server error at %s: %s", c.Path(), err.Error())

	return JSON(c, http.StatusInternalServerError, fiber.Map{
		"error": "Server error",
	})
}

func ErrExists(c *fiber.Ctx) error {
	return JSON(c, http.StatusConflict, fiber.Map{
		"error": "Entry already exists",
	})
}

func ErrNotExist(c *fiber.Ctx) error {
	return JSON(c, http.StatusNotFound, fiber.Map{
		"error": "Entry does not exist",
	})
}

func ErrBadReq(c *fiber.Ctx) error {
	return JSON(c, http.StatusBadRequest, fiber.Map{
		"error": "Provided data is invalid",
	})
}

func ErrNotFound(c *fiber.Ctx) error {
	return JSON(c, http.StatusNotFound, fiber.Map{
		"error": "Endpoint not found",
	})
}

func ErrBadJSON(c *fiber.Ctx) error {
	return JSON(c, http.StatusBadRequest, fiber.Map{
		"error": "Invalid JSON data",
	})
}

func ErrAuth(c *fiber.Ctx) error {
	return JSON(c, http.StatusUnauthorized, fiber.Map{
		"error": "Authentication failed",
	})
}

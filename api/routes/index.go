package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/config"
)

func GET_Index(c *fiber.Ctx) error {
	conf := c.Locals("config").(*config.Type)
	app := conf.GetURL("app_url")

  // redirect to the API documentation
	return c.Redirect(app.JoinPath("/doc/api").String())
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/config"
)

func GET_Index(c *fiber.Ctx) error {
	conf := c.Locals("config").(*config.Type)
	doc := conf.GetURL("doc_url")

	return c.Redirect(doc.JoinPath("/api").String())
}

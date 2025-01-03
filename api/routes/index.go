package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/config"
	"github.com/ngn13/website/api/util"
)

func GET_Index(c *fiber.Ctx) error {
	var (
		md  []byte
		err error
	)

	conf := c.Locals("config").(*config.Type)

	if !conf.GetBool("index") {
		return util.ErrNotFound(c)
	}

	frontend := conf.GetURL("frontend_url")
	api := conf.GetURL("api_url")

	if md, err = util.Render("views/index.md", fiber.Map{
		"frontend": frontend,
		"api":      api,
	}); err != nil {
		return util.ErrInternal(c, err)
	}

	return util.Markdown(c, md)
}

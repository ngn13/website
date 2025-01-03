package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/config"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

func GET_News(c *fiber.Ctx) error {
	var (
		news []database.News
		n    database.News
		feed []byte
		err  error
	)

	db := c.Locals("database").(*database.Type)
	conf := c.Locals("config").(*config.Type)
	frontend := conf.GetURL("frontend_url")
	lang := c.Params("lang")

	if lang == "" || len(lang) != 2 {
		return util.ErrBadReq(c)
	}

	lang = strings.ToLower(lang)

	for db.NewsNext(&n) {
		if n.Supports(lang) {
			news = append(news, n)
		}
	}

	if feed, err = util.Render("views/news.xml", fiber.Map{
		"frontend": frontend,
		"lang":     lang,
		"news":     news,
	}); err != nil {
		return util.ErrInternal(c, err)
	}

	c.Set("Content-Type", "application/atom+xml; charset=utf-8")
	return c.Send(feed)
}

package routes

import (
	"sort"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/config"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

// feed_entry is a temporary struct used to pass the news to the news.xml
type feed_entry struct {
	Title   string
	Author  string
	Time    time.Time
	RFC3339 string
	Content string
}

// convert UNIX timestamp to RFC3339 (format used by atom feeds)
func (e *feed_entry) From(news *database.News, lang string) {
	e.Title = news.Title.Get(lang)
	e.Author = news.Author
	e.Time = time.Unix(int64(news.Time), 0)
	e.RFC3339 = e.Time.Format(time.RFC3339)
	e.Content = news.Content.Get(lang)
}

func GET_News(c *fiber.Ctx) error {
	var (
		entries []feed_entry
		news    database.News
		indx    uint64
		feed    []byte
		err     error
	)

	db := c.Locals("database").(*database.Type)
	conf := c.Locals("config").(*config.Type)
	frontend := conf.GetURL("frontend_url")
	lang := c.Params("lang")

	if lang == "" || len(lang) != 2 {
		return util.ErrBadReq(c)
	}

	lang = strings.ToLower(lang)
	indx = 0

	for db.NewsNext(&news) {
		if news.Supports(lang) {
			entries = append(entries, feed_entry{})
			entries[indx].From(&news, lang)
			indx++
		}
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Time.Before(entries[j].Time)
	})

	if feed, err = util.Render("views/news.xml", fiber.Map{
		"frontend": frontend,
		"updated":  time.Now().Format(time.RFC3339),
		"entries":  entries,
		"lang":     lang,
	}); err != nil {
		return util.ErrInternal(c, err)
	}

	c.Set("Content-Type", "application/atom+xml; charset=utf-8")
	return c.Send(feed)
}

package routes

import (
	"database/sql"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/feeds"
	"github.com/ngn13/website/api/config"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

func GET_Post(c *fiber.Ctx) error {
	var (
		post  database.Post
		id    string
		db    *sql.DB
		found bool
		err   error
	)

	db = *(c.Locals("database").(**sql.DB))

	if id = c.Query("id"); id == "" {
		return util.ErrBadData(c)
	}

	if found, err = post.Get(db, id); err != nil {
		util.Fail("error while search for a post (\"%s\"): %s", id, err.Error())
		return util.ErrServer(c)
	}

	if !found {
		return util.ErrEntryNotExists(c)
	}

	return c.JSON(fiber.Map{
		"error":  "",
		"result": post,
	})
}

func GET_PostSum(c *fiber.Ctx) error {
	var (
		posts []database.Post
		rows  *sql.Rows
		db    *sql.DB
		err   error
	)

	db = *(c.Locals("database").(**sql.DB))

	if rows, err = db.Query("SELECT * FROM posts"); err != nil {
		util.Fail("cannot load posts: %s", err.Error())
		return util.ErrServer(c)
	}
	defer rows.Close()

	for rows.Next() {
		var post database.Post

		if err = post.Load(rows); err != nil {
			util.Fail("error while loading post: %s", err.Error())
			return util.ErrServer(c)
		}

		if post.Public == 0 {
			continue
		}

		if len(post.Content) > 255 {
			post.Content = post.Content[0:250]
		}

		posts = append(posts, post)
	}

	return c.JSON(fiber.Map{
		"error":  "",
		"result": posts,
	})
}

func getFeed(db *sql.DB) (*feeds.Feed, error) {
	var (
		posts []database.Post
		err   error
	)

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post database.Post

		if err = post.Load(rows); err != nil {
			return nil, err
		}

		if post.Public == 0 {
			continue
		}

		posts = append(posts, post)
	}
	rows.Close()

	blogurl, err := url.JoinPath(
		config.Get("frontend_url"), "/blog",
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create the blog URL: %s", err.Error())
	}

	feed := &feeds.Feed{
		Title:       "[ngn.tf] | blog",
		Link:        &feeds.Link{Href: blogurl},
		Description: "ngn's personal blog",
		Author:      &feeds.Author{Name: "ngn", Email: "ngn@ngn.tf"},
		Created:     time.Now(),
	}

	feed.Items = []*feeds.Item{}
	for _, p := range posts {
		purl, err := url.JoinPath(blogurl, p.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to create URL for '%s': %s\n", p.ID, err.Error())
		}

		parsed, err := time.Parse("02/01/06", p.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to parse time for '%s': %s\n", p.ID, err.Error())
		}

		feed.Items = append(feed.Items, &feeds.Item{
			Id:      p.ID,
			Title:   p.Title,
			Link:    &feeds.Link{Href: purl},
			Author:  &feeds.Author{Name: p.Author},
			Created: parsed,
		})
	}

	return feed, nil
}

func GET_Feed(c *fiber.Ctx) error {
	var (
		db   *sql.DB
		err  error
		feed *feeds.Feed
		name []string
		res  string
		ext  string
	)

	db = *(c.Locals("database").(**sql.DB))

	if name = strings.Split(path.Base(c.Path()), "."); len(name) != 2 {
		return util.ErrNotFound(c)
	}
	ext = name[1]

	if feed, err = getFeed(db); err != nil {
		util.Fail("cannot obtain the feed: %s", err.Error())
		return util.ErrServer(c)
	}

	switch ext {
	case "atom":
		res, err = feed.ToAtom()
		c.Set("Content-Type", "application/atom+xml")
		break

	case "json":
		res, err = feed.ToJSON()
		c.Set("Content-Type", "application/feed+json")
		break

	case "rss":
		res, err = feed.ToRss()
		c.Set("Content-Type", "application/rss+xml")
		break

	default:
		return util.ErrNotFound(c)
	}

	if err != nil {
		util.Fail("cannot obtain the feed as the specified format: %s", err.Error())
		return util.ErrServer(c)
	}

	return c.Send([]byte(res))
}

package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/feeds"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/global"
	"github.com/ngn13/website/api/util"
)

func PostFromRow(post *global.Post, rows *sql.Rows) error {
	err := rows.Scan(&post.ID, &post.Title, &post.Author, &post.Date, &post.Content, &post.Public, &post.Vote)
	if err != nil {
		return err
	}

	return nil
}

func GetPostByID(db *database.Type, id string) (global.Post, string) {
	var post global.Post = global.Post{}
	post.Title = "NONE"

	rows, err := db.Sql.Query("SELECT * FROM posts WHERE id = ?", id)

	if err != nil {
		return post, "Server error"
	}

	success := rows.Next()
	if !success {
		rows.Close()
		return post, "Post not found"
	}

	err = PostFromRow(&post, rows)
	if err != nil {
		rows.Close()
		return post, "Server error"
	}
	rows.Close()

	if post.Title == "NONE" {
		return post, "Post not found"
	}

	return post, ""
}

func VoteStat(c *fiber.Ctx) error {
	var (
		db *database.Type
		id string
	)

	db = c.Locals("database").(*database.Type)
	id = c.Query("id")

	if id == "" {
		return util.ErrBadData(c)
	}

	for i := 0; i < len(db.Votes); i++ {
		if db.Votes[i].Client == util.GetIP(c) && db.Votes[i].Post == id {
			return c.JSON(fiber.Map{
				"error":  "",
				"result": db.Votes[i].Status,
			})
		}
	}

	return c.Status(http.StatusNotFound).JSON(util.ErrorJSON("Client never voted"))
}

func VoteSet(c *fiber.Ctx) error {
	var (
		id    string
		to    string
		voted bool
		db    *database.Type
	)

	db = c.Locals("database").(*database.Type)
	id = c.Query("id")
	to = c.Query("to")
	voted = false

	if id == "" || (to != "upvote" && to != "downvote") {
		return util.ErrBadData(c)
	}

	for i := 0; i < len(db.Votes); i++ {
		if db.Votes[i].Client == util.GetIP(c) && db.Votes[i].Post == id && db.Votes[i].Status == to {
			return c.Status(http.StatusForbidden).JSON(util.ErrorJSON("Client already voted"))
		}

		if db.Votes[i].Client == util.GetIP(c) && db.Votes[i].Post == id && db.Votes[i].Status != to {
			voted = true
		}
	}

	post, msg := GetPostByID(db, id)
	if msg != "" {
		return c.Status(http.StatusNotFound).JSON(util.ErrorJSON(msg))
	}

	vote := post.Vote + 1

	if to == "downvote" {
		vote = post.Vote - 1
	}

	if to == "downvote" && voted {
		vote = vote - 1
	}

	if to == "upvote" && voted {
		vote = vote + 1
	}

	_, err := db.Sql.Exec("UPDATE posts SET vote = ? WHERE title = ?", vote, post.Title)
	if util.ErrorCheck(err, c) {
		return util.ErrServer(c)
	}

	for i := 0; i < len(db.Votes); i++ {
		if db.Votes[i].Client == util.GetIP(c) && db.Votes[i].Post == id && db.Votes[i].Status != to {
			db.Votes[i].Status = to
			return util.NoError(c)
		}
	}

	var entry = global.Vote{}
	entry.Client = util.GetIP(c)
	entry.Status = to
	entry.Post = id
	db.Votes = append(db.Votes, entry)
	return util.NoError(c)
}

func GetPost(c *fiber.Ctx) error {
	var (
		id string
		db *database.Type
	)

	id = c.Query("id")
	db = c.Locals("database").(*database.Type)

	if id == "" {
		return util.ErrBadData(c)
	}

	post, msg := GetPostByID(db, id)
	if msg != "" {
		return c.Status(http.StatusNotFound).JSON(util.ErrorJSON(msg))
	}

	return c.JSON(fiber.Map{
		"error":  "",
		"result": post,
	})
}

func SumPost(c *fiber.Ctx) error {
	var (
		posts []global.Post
		post  global.Post
		db    *database.Type
		err   error
	)

	db = c.Locals("database").(*database.Type)

	rows, err := db.Sql.Query("SELECT * FROM posts")
	if util.ErrorCheck(err, c) {
		return util.ErrServer(c)
	}

	for rows.Next() {
		err = PostFromRow(&post, rows)

		if util.ErrorCheck(err, c) {
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
	rows.Close()

	return c.JSON(fiber.Map{
		"error":  "",
		"result": posts,
	})
}

func GetFeed(db *database.Type) (*feeds.Feed, error) {
	var (
		posts []global.Post
		post  global.Post
		err   error
	)

	rows, err := db.Sql.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := PostFromRow(&post, rows)

		if err != nil {
			return nil, err
		}

		if post.Public == 0 {
			continue
		}

		posts = append(posts, post)
	}
	rows.Close()

	blogurl, err := url.JoinPath(os.Getenv("FRONTEND_URL"), "/blog")
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

func GetAtomFeed(c *fiber.Ctx) error {
	feed, err := GetFeed(c.Locals("database").(*database.Type))
	if util.ErrorCheck(err, c) {
		return util.ErrServer(c)
	}

	atom, err := feed.ToAtom()
	if util.ErrorCheck(err, c) {
		return util.ErrServer(c)
	}

	c.Set("Content-Type", "application/atom+xml")
	return c.Send([]byte(atom))
}

func GetRSSFeed(c *fiber.Ctx) error {
	feed, err := GetFeed(c.Locals("database").(*database.Type))
	if util.ErrorCheck(err, c) {
		return util.ErrServer(c)
	}

	rss, err := feed.ToRss()
	if util.ErrorCheck(err, c) {
		return util.ErrServer(c)
	}

	c.Set("Content-Type", "application/rss+xml")
	return c.Send([]byte(rss))
}

func GetJSONFeed(c *fiber.Ctx) error {
	feed, err := GetFeed(c.Locals("database").(*database.Type))
	if util.ErrorCheck(err, c) {
		return util.ErrServer(c)
	}

	json, err := feed.ToJSON()
	if util.ErrorCheck(err, c) {
		return util.ErrServer(c)
	}
	c.Set("Content-Type", "application/feed+json")
	return c.Send([]byte(json))
}

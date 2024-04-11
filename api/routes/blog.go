package routes

import (
	"database/sql"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/feeds"
	"github.com/ngn13/website/api/util"
)

func BlogDb(db *sql.DB) {
  _, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS posts(
      id      TEXT NOT NULL UNIQUE,
      title   TEXT NOT NULL,
      author  TEXT NOT NULL,
      date    TEXT NOT NULL,
      content TEXT NOT NULL,
      public  INTEGER NOT NULL,
      vote    INTEGER NOT NULL
    );
  `)
  DB = db
  if err != nil {
    log.Fatal("Error creating table: "+err.Error())
  }
}

func GetIP(c *fiber.Ctx) string {
  if c.Get("X-Real-IP") != "" {
    return strings.Clone(c.Get("X-Real-IP"))
  }

  return c.IP()
}

func VoteStat(c *fiber.Ctx) error{
  var id = c.Query("id")
  if id == "" {
    return util.ErrBadData(c)
  }

  for i := 0; i < len(votelist); i++ {
    if votelist[i].Client == GetIP(c) && votelist[i].Post == id {
      return c.JSON(fiber.Map {
        "error": "",
        "result": votelist[i].Status,
      })
    }
  }

  return c.Status(http.StatusNotFound).JSON(util.ErrorJSON("Client never voted"))
}

func VoteSet(c *fiber.Ctx) error{
  var id = c.Query("id")
  var to = c.Query("to")
  voted := false

  if id == "" || (to != "upvote" && to != "downvote") {
    return util.ErrBadData(c)
  }

  for i := 0; i < len(votelist); i++ {
    if votelist[i].Client == GetIP(c) && votelist[i].Post == id && votelist[i].Status == to {
      return c.Status(http.StatusForbidden).JSON(util.ErrorJSON("Client already voted"))
    }

    if votelist[i].Client == GetIP(c) && votelist[i].Post == id && votelist[i].Status != to {
      voted = true
    }
  }

  post, msg := GetPostByID(id)
  if msg != ""{
    return c.Status(http.StatusNotFound).JSON(util.ErrorJSON(msg))
  }

  vote := post.Vote+1

  if to == "downvote" {
    vote = post.Vote-1
  }

  if to == "downvote" && voted {
    vote = vote-1
  } 

  if to == "upvote" && voted {
    vote = vote+1
  }

  _, err := DB.Exec("UPDATE posts SET vote = ? WHERE title = ?", vote, post.Title)
  if util.ErrorCheck(err, c){
    return util.ErrServer(c)
  }

  for i := 0; i < len(votelist); i++ {
    if votelist[i].Client == GetIP(c) && votelist[i].Post == id && votelist[i].Status != to {
      votelist[i].Status = to
      return util.NoError(c)
    }
  }

  var entry = Vote{}
  entry.Client = GetIP(c)
  entry.Status = to 
  entry.Post = id 
  votelist = append(votelist, entry)
  return util.NoError(c)
}

func GetPost(c *fiber.Ctx) error{
  var id = c.Query("id")  
  if id == "" {
    return util.ErrBadData(c)
  }

  post, msg := GetPostByID(id)
  if msg != ""{
    return c.Status(http.StatusNotFound).JSON(util.ErrorJSON(msg))
  } 

  return c.JSON(fiber.Map {
    "error": "",
    "result": post, 
  })
}

func GetFeed(c *fiber.Ctx) error{
  var posts []Post = []Post{}
  rows, err := DB.Query("SELECT * FROM posts")
  if util.ErrorCheck(err, c) {
    return util.ErrServer(c)
  } 

  for rows.Next() {
    var post Post
    err := PostFromRow(&post, rows)

    if util.ErrorCheck(err, c) {
      return util.ErrServer(c)
    }
    
    if post.Public == 0 {
      continue
    }

    posts = append(posts, post)
  }
  rows.Close()


  blogurl, err := url.JoinPath(os.Getenv("URL"), "/blog")
  if err != nil {
    log.Printf("Failed to create the blog URL: %s\n", err.Error())
    return c.JSON(fiber.Map{"error": "Server error"})
  }

  feed := &feeds.Feed{
    Title:       "[ngn] | blog",
    Link:        &feeds.Link{Href: blogurl},
    Description: "ngn's personal blog",
    Author:      &feeds.Author{Name: "ngn", Email: "ngn@ngn.tf"},
    Created:     time.Now(),
  }

  feed.Items = []*feeds.Item{}
  for _, p := range posts {
    purl, err := url.JoinPath(blogurl, p.ID)
    if err != nil {
      log.Printf("Failed to create URL for '%s': %s\n", p.ID, err.Error())
      continue
    }

    parsed, err := time.Parse("02/01/06", p.Date)
    if err != nil {
      log.Printf("Failed to parse time for '%s': %s\n", p.ID, err.Error())
      continue
    }

    feed.Items = append(feed.Items, &feeds.Item{
      Title:   p.Title,
      Link:    &feeds.Link{Href: purl},
      Author:  &feeds.Author{Name: p.Author},
      Created: parsed,
    })
  }

  atom, err := feed.ToAtom()
  if err != nil {
    log.Printf("Failed to create atom feed: %s", err.Error())
    return c.JSON(fiber.Map{"error": "Server error"})
  }

  c.Set("Content-Type", "application/atom+xml")
  return c.Send([]byte(atom))
}

func SumPost(c *fiber.Ctx) error{
  var posts []Post = []Post{}
  rows, err := DB.Query("SELECT * FROM posts")
  if util.ErrorCheck(err, c) {
    return util.ErrServer(c)
  } 

  for rows.Next() {
    var post Post
    err := PostFromRow(&post, rows)

    if util.ErrorCheck(err, c) {
      return util.ErrServer(c)
    }
    
    if post.Public == 0 {
      continue
    }

    if len(post.Content) > 255{
      post.Content = post.Content[0:250]
    }

    posts = append(posts, post)
  }
  rows.Close()

  return c.JSON(fiber.Map {
    "error": "",
    "result": posts, 
  })
}

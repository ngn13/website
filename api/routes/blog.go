package routes

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
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

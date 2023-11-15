package routes

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-sqlite3"
	"github.com/ngn13/website/api/util"
)

var Token string = util.CreateToken() 

func AuthMiddleware(c *fiber.Ctx) error {
  if c.Path() == "/admin/login" {
    return c.Next()
  }

  if c.Get("Authorization") != Token {
    return util.ErrAuth(c) 
  }

  return c.Next()
}

func Login(c *fiber.Ctx) error{
  if c.Query("pass") != os.Getenv("PASSWORD") {
    return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
      "error": "Authentication failed", 
    })
  }

  return c.Status(http.StatusOK).JSON(fiber.Map{
    "error": "",
    "token": Token,
  })
}

func Logout(c *fiber.Ctx) error{
  Token = util.CreateToken()
  return c.Status(http.StatusOK).JSON(fiber.Map{
    "error": "",
  })
}

func RemoveService(c *fiber.Ctx) error {
  name := c.Query("name")
  if name == "" {
    util.ErrBadData(c)
  }

  _, err := DB.Exec("DELETE FROM services WHERE name = ?", name)
  if util.ErrorCheck(err, c){
    return util.ErrServer(c)
  }
 
  return util.NoError(c)
}

func AddService(c *fiber.Ctx) error {
  var service Service 
  if c.BodyParser(&service) != nil {
    return util.ErrBadJSON(c)
  }

  if service.Name == "" || service.Desc == "" || service.Url == "" {
    return util.ErrBadData(c)
  }

  rows, err := DB.Query("SELECT * FROM services WHERE name = ?", service.Name)
  if util.ErrorCheck(err, c){
    return util.ErrServer(c)
  }
 
  if rows.Next() {
    rows.Close()
    return util.ErrExists(c)
  }

  rows.Close()

  _, err = DB.Exec(
    "INSERT INTO services(name, desc, url) values(?, ?, ?)",
    service.Name, service.Desc, service.Url,
  )

  if util.ErrorCheck(err, c){
    return util.ErrServer(c)
  }

  return util.NoError(c)
}

func RemovePost(c *fiber.Ctx) error{
  var id = c.Query("id")
  if id == "" {
    return util.ErrBadData(c)
  }

  _, err := DB.Exec("DELETE FROM posts WHERE id = ?", id)
  if util.ErrorCheck(err, c){
    return util.ErrServer(c)
  }

  return util.NoError(c)
}

func AddPost(c *fiber.Ctx) error{
  var post Post
  post.Public = 1

  if c.BodyParser(&post) != nil {
    return util.ErrBadJSON(c)
  }

  if post.Title == "" || post.Author == "" || post.Content == "" {
    return util.ErrBadData(c)
  }

  post.Date = time.Now().Format("02/01/06")
  post.ID = TitleToID(post.Title)

  _, err := DB.Exec(
    "INSERT INTO posts(id, title, author, date, content, public, vote) values(?, ?, ?, ?, ?, ?, ?)",
    post.ID, post.Title, post.Author, post.Date, post.Content, post.Public, post.Vote,
  )

  if err != nil && strings.Contains(err.Error(), sqlite3.ErrConstraintUnique.Error()) {
    return util.ErrExists(c)
  }

  if util.ErrorCheck(err, c){
    return util.ErrExists(c)
  }

  return util.NoError(c)
}

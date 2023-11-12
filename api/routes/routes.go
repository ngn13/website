package routes

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var DB *sql.DB

func Setup(app *fiber.App, db *sql.DB){
  // database init
  DB = db

  // index route 
  app.Get("/", func(c *fiber.Ctx) error {
    return c.Send([]byte("o/"))
  })

  // blog routes 
  app.Get("/blog/sum", SumPost)
  app.Get("/blog/get", GetPost)
  app.Get("/blog/vote/set", VoteSet)
  app.Get("/blog/vote/status", VoteStat)

  // service routes 
  app.Get("/services/all", GetServices)

  // admin routes 
  app.Use("/admin*", AuthMiddleware)
  app.Get("/admin/login", Login)
  app.Get("/admin/logout", Logout)
  app.Put("/admin/service/add", AddService)
  app.Delete("/admin/service/remove", RemoveService)
  app.Put("/admin/blog/add", AddPost)
  app.Delete("/admin/blog/remove", RemovePost)

  // 404 page 
  app.All("*", func(c *fiber.Ctx) error {
    return c.Status(http.StatusNotFound).JSON(fiber.Map {
      "error": "Requested endpoint not found",
    })
  })
}

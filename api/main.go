package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/routes"
)

func CorsMiddleware(c *fiber.Ctx) error {
  c.Set("Access-Control-Allow-Origin", "*")
  c.Set("Access-Control-Allow-Headers", 
  "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
  c.Set("Access-Control-Allow-Methods", "PUT, DELETE, GET")
  return c.Next()
}

func main() {
  app := fiber.New(fiber.Config{
    DisableStartupMessage: true,
  })
  app.Use(CorsMiddleware)

  db, err := sql.Open("sqlite3", "data.db")
  if err != nil {
    log.Fatal("Cannot connect to the database: "+err.Error())
  }

  log.Println("Creating tables")
  routes.BlogDb(db)
  routes.ServicesDb(db)
  routes.Setup(app, db)

  log.Println("Starting web server at port 7001")
  err = app.Listen("0.0.0.0:7001")
  if err != nil {
    log.Printf("Error starting the webserver: %s", err.Error())
  }

  defer db.Close()
}

package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/routes"
)

func main() {
	var (
		app *fiber.App
		db  database.Type
		err error
	)

	if err = db.Open("data.db"); err != nil {
		log.Fatalf("Cannot access the database: %s", err.Error())
		return
	}
	defer db.Close()

	app = fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use("*", func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Set("Access-Control-Allow-Methods", "PUT, DELETE, GET")
		c.Locals("database", &db)
		return c.Next()
	})

	// index route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("o/"))
	})

	// blog routes
	blog_routes := app.Group("/blog")
	blog_routes.Get("/feed.atom", routes.GetAtomFeed)
	blog_routes.Get("/feed.rss", routes.GetRSSFeed)
	blog_routes.Get("/feed.json", routes.GetJSONFeed)
	blog_routes.Get("/sum", routes.SumPost)
	blog_routes.Get("/get", routes.GetPost)
	blog_routes.Get("/vote/set", routes.VoteSet)
	blog_routes.Get("/vote/status", routes.VoteStat)

	// service routes
	service_routes := app.Group("services")
	service_routes.Get("/all", routes.GetServices)

	// admin routes
	admin_routes := app.Group("admin")
	admin_routes.Use("*", routes.AuthMiddleware)
	admin_routes.Get("/login", routes.Login)
	admin_routes.Get("/logout", routes.Logout)
	admin_routes.Put("/service/add", routes.AddService)
	admin_routes.Delete("/service/remove", routes.RemoveService)
	admin_routes.Put("/blog/add", routes.AddPost)
	admin_routes.Delete("/blog/remove", routes.RemovePost)

	// 404 routes
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Requested endpoint not found",
		})
	})

	log.Println("Starting web server at port 7001")
	if err = app.Listen("0.0.0.0:7001"); err != nil {
		log.Printf("Error starting the webserver: %s", err.Error())
	}
}

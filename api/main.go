package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/config"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/routes"
	"github.com/ngn13/website/api/util"
)

var db *sql.DB

func main() {
	var (
		app *fiber.App
		//db  *sql.DB
		err error
	)

	if !config.Load() {
		util.Fail("failed to load the configuration")
		return
	}

	if db, err = sql.Open("sqlite3", "data.db"); err != nil {
		util.Fail("cannot access the database: %s", err.Error())
		return
	}
	defer db.Close()

	if err = database.Setup(db); err != nil {
		util.Fail("cannot setup the database: %s", err.Error())
		return
	}

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

	// blog feed routes
	blog_routes.Get("/feed.*", routes.GET_Feed)

	// blog post routes
	blog_routes.Get("/sum", routes.GET_PostSum)
	blog_routes.Get("/get", routes.GET_Post)

	// blog post vote routes
	blog_routes.Get("/vote/set", routes.GET_VoteSet)
	blog_routes.Get("/vote/get", routes.GET_VoteGet)

	// service routes
	service_routes := app.Group("services")
	service_routes.Get("/all", routes.GET_Services)

	// admin routes
	admin_routes := app.Group("admin")
	admin_routes.Use("*", routes.AuthMiddleware)

	// admin auth routes
	admin_routes.Get("/login", routes.GET_Login)
	admin_routes.Get("/logout", routes.GET_Logout)

	// admin service managment routes
	admin_routes.Put("/service/add", routes.PUT_AddService)
	admin_routes.Delete("/service/remove", routes.DEL_RemoveService)

	// admin blog managment routes
	admin_routes.Put("/blog/add", routes.PUT_AddPost)
	admin_routes.Delete("/blog/remove", routes.DEL_RemovePost)

	// 404 route
	app.All("*", func(c *fiber.Ctx) error {
		return util.ErrNotFound(c)
	})

	util.Info("starting web server at port 7001")

	if err = app.Listen("0.0.0.0:7001"); err != nil {
		util.Fail("error starting the webserver: %s", err.Error())
	}
}

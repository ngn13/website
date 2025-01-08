package main

/*

 * website/api | API server for my personal website
 * written by ngn (https://ngn.tf) (2025)

 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.

 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.

 */

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/config"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/routes"
	"github.com/ngn13/website/api/status"
	"github.com/ngn13/website/api/util"
)

func main() {
	var (
		app  *fiber.App
		stat status.Type

		conf config.Type
		db   database.Type

		err error
	)

	if err = conf.Load(); err != nil {
		util.Fail("failed to load the configuration: %s", err.Error())
		return
	}

	if !conf.GetBool("debug") {
		util.Debg = func(m string, v ...any) {}
	}

	if err = db.Load(); err != nil {
		util.Fail("failed to load the database: %s", err.Error())
		return
	}

	if err = stat.Setup(&conf, &db); err != nil {
		util.Fail("failed to setup the status checker: %s", err.Error())
		return
	}

	app = fiber.New(fiber.Config{
		AppName:               "ngn's website",
		DisableStartupMessage: true,
		ServerHeader:          "",
	})

	app.Use("*", func(c *fiber.Ctx) error {
		// CORS stuff
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Set("Access-Control-Allow-Methods", "PUT, DELETE, GET") // POST can be sent from HTML forms, so I prefer PUT for API endpoints

		c.Locals("status", &stat)
		c.Locals("config", &conf)
		c.Locals("database", &db)

		return c.Next()
	})

	// index route
	app.Get("/", routes.GET_Index)

	// version groups
	v1 := app.Group("v1")

	// v1 user routes
	v1.Get("/services", routes.GET_Services)
	v1.Get("/projects", routes.GET_Projects)
	v1.Get("/metrics", routes.GET_Metrics)
	v1.Get("/news/:lang", routes.GET_News)

	// v1 admin routes
	v1.Use("/admin", routes.AuthMiddleware)
	v1.Get("/admin/logs", routes.GET_AdminLogs)

	v1.Get("/admin/service/check", routes.GET_CheckService)
	v1.Put("/admin/service/add", routes.PUT_AddService)
	v1.Delete("/admin/service/del", routes.DEL_DelService)

	v1.Put("/admin/project/add", routes.PUT_AddProject)
	v1.Delete("/admin/project/del", routes.DEL_DelProject)

	v1.Put("/admin/news/add", routes.PUT_AddNews)
	v1.Delete("/admin/news/del", routes.DEL_DelNews)

	// 404 route
	app.All("*", func(c *fiber.Ctx) error {
		return util.JSON(c, http.StatusNotFound, fiber.Map{
			"error": "Endpoint not found",
		})
	})

	// start the status checker
	if err = stat.Run(); err != nil {
		util.Fail("failed to start the status checker: %s", err.Error())
		return
	}

	// start the app
	util.Info("starting web server on %s", conf.GetStr("host"))

	if err = app.Listen(conf.GetStr("host")); err != nil {
		util.Fail("failed to start the web server: %s", err.Error())
	}

	stat.Stop()
}

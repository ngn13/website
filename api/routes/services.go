package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

func GET_Services(c *fiber.Ctx) error {
	var (
		services []database.Service
		rows     *sql.Rows
		db       *sql.DB
		err      error
	)

	db = *(c.Locals("database").(**sql.DB))

	if rows, err = db.Query("SELECT * FROM services"); err != nil {
		util.Fail("cannot load services: %s", err.Error())
		return util.ErrServer(c)
	}
	defer rows.Close()

	for rows.Next() {
		var service database.Service

		if err = service.Load(rows); err != nil {
			util.Fail("error while loading service: %s", err.Error())
			return util.ErrServer(c)
		}

		services = append(services, service)
	}

	return c.JSON(fiber.Map{
		"error":  "",
		"result": services,
	})
}

package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/global"
	"github.com/ngn13/website/api/util"
)

func GetServices(c *fiber.Ctx) error {
	var (
		services []global.Service = []global.Service{}
		service  global.Service
		db       *database.Type
		err      error
	)

	db = c.Locals("database").(*database.Type)

	rows, err := db.Sql.Query("SELECT * FROM services")
	if util.ErrorCheck(err, c) {
		return util.ErrServer(c)
	}

	for rows.Next() {
		if err = rows.Scan(&service.Name, &service.Desc, &service.Url); err != nil {
			log.Println("Error scaning services row: " + err.Error())
			continue
		}
		services = append(services, service)
	}

	rows.Close()

	return c.JSON(fiber.Map{
		"error":  "",
		"result": services,
	})
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

func GET_Services(c *fiber.Ctx) error {
	var (
		services []database.Service
		service  database.Service
	)

	db := c.Locals("database").(*database.Type)
	name := c.Query("name")

	if name != "" {
		if s, err := db.ServiceFind(name); err != nil {
			return util.ErrInternal(c, err)
		} else if s != nil {
			return util.JSON(c, 200, fiber.Map{
				"result": s,
			})
		}

		return util.ErrNotExist(c)
	}

	for db.ServiceNext(&service) {
		services = append(services, service)
	}

	return util.JSON(c, 200, fiber.Map{
		"result": services,
	})
}

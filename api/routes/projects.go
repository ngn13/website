package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

func GET_Projects(c *fiber.Ctx) error {
	var (
		projects []database.Project
		project  database.Project
	)

	db := c.Locals("database").(*database.Type)

	for db.ProjectNext(&project) {
		projects = append(projects, project)
	}

	return util.JSON(c, 200, fiber.Map{
		"result": projects,
	})
}

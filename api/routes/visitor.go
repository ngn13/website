package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

const LAST_ADDRS_MAX = 30

var last_addrs []string

func GET_Visitor(c *fiber.Ctx) error {
	var (
		err   error
		count uint64
	)

	db := c.Locals("database").(*database.Type)
	new_addr := util.GetSHA1(util.IP(c))

	for _, addr := range last_addrs {
		if new_addr == addr {
			if count, err = db.VisitorGet(); err != nil {
				return util.ErrInternal(c, err)
			}

			return util.JSON(c, 200, fiber.Map{
				"result": count,
			})
		}
	}

	if err = db.VisitorIncrement(); err != nil {
		return util.ErrInternal(c, err)
	}

	if count, err = db.VisitorGet(); err != nil {
		return util.ErrInternal(c, err)
	}

	if len(last_addrs) > LAST_ADDRS_MAX {
		last_addrs = append(last_addrs[:0], last_addrs[1:]...)
		last_addrs = append(last_addrs, new_addr)
	}

	return util.JSON(c, 200, fiber.Map{
		"result": count,
	})
}

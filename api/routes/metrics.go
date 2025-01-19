package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

const VISITOR_CACHE_MAX = 30 // store 30 visitor data at most
var visitor_cache []string   // in memory cache for the visitor addresses

func GET_Metrics(c *fiber.Ctx) error {
	var (
		err    error
		result map[string]uint64 = map[string]uint64{
			"total": 0, // total number of visitors
			"since": 0, // metric collection start date (UNIX timestamp)
		}
	)

	db := c.Locals("database").(*database.Type)
	new_addr := util.GetSHA1(util.IP(c))
	is_in_cache := false

	for _, cache := range visitor_cache {
		if new_addr == cache {
			is_in_cache = true
			break
		}
	}

	if result["total"], err = db.MetricsGet("visitor_count"); err != nil {
		return util.ErrInternal(c, err)
	}

	if !is_in_cache {
		if len(visitor_cache) > VISITOR_CACHE_MAX {
			util.Debg("visitor cache is full, removing the oldest entry")
			visitor_cache = visitor_cache[1:]
		}

		visitor_cache = append(visitor_cache, new_addr)
		result["total"]++

		if err = db.MetricsSet("visitor_count", result["total"]); err != nil {
			return util.ErrInternal(c, err)
		}
	}

	if result["since"], err = db.MetricsGet("start_date"); err != nil {
		return util.ErrInternal(c, err)
	}

	if result["since"] == 0 {
		result["since"] = uint64(time.Now().Truncate(24 * time.Hour).Unix())

		if err = db.MetricsSet("start_date", result["since"]); err != nil {
			return util.ErrInternal(c, err)
		}
	}

	return util.JSON(c, 200, fiber.Map{
		"result": result,
	})
}

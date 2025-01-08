package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

type visitor_cache_entry struct {
	Addr   string // SHA1 hash of visitor's IP
	Number uint64 // number of the visitor
}

const VISITOR_CACHE_MAX = 30            // store 30 visitor data at most
var visitor_cache []visitor_cache_entry // in memory cache for the visitor

func GET_Metrics(c *fiber.Ctx) error {
	var (
		err    error
		result map[string]uint64 = map[string]uint64{
			"number": 0, // visitor number of the current visitor
			"total":  0, // total number of visitors
			"since":  0, // metric collection start date (UNIX timestamp)
		}
	)

	db := c.Locals("database").(*database.Type)
	new_addr := util.GetSHA1(util.IP(c))

	for i := range visitor_cache {
		if new_addr == visitor_cache[i].Addr {
			result["number"] = visitor_cache[i].Number
			break
		}
	}

	if result["total"], err = db.MetricsGet("visitor_count"); err != nil {
		return util.ErrInternal(c, err)
	}

	if result["number"] == 0 {
		result["total"]++
		result["number"] = result["total"]

		if len(visitor_cache) > VISITOR_CACHE_MAX {
			util.Debg("visitor cache is full, removing the oldest entry")
			visitor_cache = visitor_cache[1:]
		}

		visitor_cache = append(visitor_cache, visitor_cache_entry{
			Addr:   new_addr,
			Number: result["number"],
		})

		if err = db.MetricsSet("visitor_count", result["total"]); err != nil {
			return util.ErrInternal(c, err)
		}
	}

	if result["since"], err = db.MetricsGet("start_date"); err != nil {
		return util.ErrInternal(c, err)
	}

	if result["since"] == 0 {
		result["since"] = uint64(time.Now().Truncate(24 * time.Hour).Unix())

		if err = db.MetricsSet("since", result["since"]); err != nil {
			return util.ErrInternal(c, err)
		}
	}

	return util.JSON(c, 200, fiber.Map{
		"result": result,
	})
}

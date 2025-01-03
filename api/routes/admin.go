package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/config"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/status"
	"github.com/ngn13/website/api/util"
)

func admin_log(c *fiber.Ctx, m string) error {
	return c.Locals("database").(*database.Type).AdminLogAdd(&database.AdminLog{
		Action: m,                 // action that the admin peformed
		Time:   time.Now().Unix(), // current time
	})
}

func AuthMiddleware(c *fiber.Ctx) error {
	conf := c.Locals("config").(*config.Type)

	if c.Get("Authorization") != conf.GetStr("password") {
		return util.ErrAuth(c)
	}

	return c.Next()
}

func GET_AdminLogs(c *fiber.Ctx) error {
	var (
		list []database.AdminLog
		log  database.AdminLog
	)

	db := c.Locals("database").(*database.Type)

	for db.AdminLogNext(&log) {
		list = append(list, log)
	}

	return util.JSON(c, 200, fiber.Map{
		"result": list,
	})
}

func DEL_DelService(c *fiber.Ctx) error {
	var (
		name string
		err  error
	)

	db := c.Locals("database").(*database.Type)

	if name = c.Query("name"); name == "" {
		util.ErrBadReq(c)
	}

	if err = admin_log(c, fmt.Sprintf("Removed service \"%s\"", name)); err != nil {
		return util.ErrInternal(c, err)
	}

	if err = db.ServiceRemove(name); err != nil {
		return util.ErrInternal(c, err)
	}

	return util.JSON(c, 200, nil)
}

func PUT_AddService(c *fiber.Ctx) error {
	var (
		service database.Service
		err     error
	)

	db := c.Locals("database").(*database.Type)

	if c.BodyParser(&service) != nil {
		return util.ErrBadJSON(c)
	}

	if !service.IsValid() {
		return util.ErrBadReq(c)
	}

	if err = admin_log(c, fmt.Sprintf("Added service \"%s\"", service.Name)); err != nil {
		return util.ErrInternal(c, err)
	}

	if err = db.ServiceUpdate(&service); err != nil {
		return util.ErrInternal(c, err)
	}

	// force a status check so we can get the status of the new service
	c.Locals("status").(*status.Type).Check()

	return util.JSON(c, 200, nil)
}

func GET_CheckService(c *fiber.Ctx) error {
	c.Locals("status").(*status.Type).Check()
	return util.JSON(c, 200, nil)
}

func DEL_DelNews(c *fiber.Ctx) error {
	var (
		id  string
		err error
	)

	db := c.Locals("database").(*database.Type)

	if id = c.Query("id"); id == "" {
		util.ErrBadReq(c)
	}

	if err = admin_log(c, fmt.Sprintf("Removed news \"%s\"", id)); err != nil {
		return util.ErrInternal(c, err)
	}

	if err = db.NewsRemove(id); err != nil {
		return util.ErrInternal(c, err)
	}

	return util.JSON(c, 200, nil)
}

func PUT_AddNews(c *fiber.Ctx) error {
	var (
		news database.News
		err  error
	)

	db := c.Locals("database").(*database.Type)

	if c.BodyParser(&news) != nil {
		return util.ErrBadJSON(c)
	}

	if !news.IsValid() {
		return util.ErrBadReq(c)
	}

	if err = admin_log(c, fmt.Sprintf("Added news \"%s\"", news.ID)); err != nil {
		return util.ErrInternal(c, err)
	}

	if err = db.NewsAdd(&news); err != nil {
		return util.ErrInternal(c, err)
	}

	return util.JSON(c, 200, nil)
}

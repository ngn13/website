package routes

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-sqlite3"
	"github.com/ngn13/website/api/config"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

var Token string = util.CreateToken()

func AuthMiddleware(c *fiber.Ctx) error {
	if c.Path() == "/admin/login" {
		return c.Next()
	}

	if c.Get("Authorization") != Token {
		return util.ErrAuth(c)
	}

	return c.Next()
}

func GET_Login(c *fiber.Ctx) error {
	if c.Query("pass") != config.Get("password") {
		return util.ErrAuth(c)
	}

	util.Info("new login from %s", util.GetIP(c))

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"error": "",
		"token": Token,
	})
}

func GET_Logout(c *fiber.Ctx) error {
	Token = util.CreateToken()

	util.Info("logout from %s", util.GetIP(c))

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"error": "",
	})
}

func DEL_RemoveService(c *fiber.Ctx) error {
	var (
		db      *sql.DB
		service database.Service
		name    string
		found   bool
		err     error
	)

	db = *(c.Locals("database").(**sql.DB))
	name = c.Query("name")

	if name == "" {
		util.ErrBadData(c)
	}

	if found, err = service.Get(db, name); err != nil {
		util.Fail("error while searching for a service (\"%s\"): %s", name, err.Error())
		return util.ErrServer(c)
	}

	if !found {
		return util.ErrEntryNotExists(c)
	}

	if err = service.Remove(db); err != nil {
		util.Fail("error while removing a service (\"%s\"): %s", service.Name, err.Error())
		return util.ErrServer(c)
	}

	return util.NoError(c)
}

func PUT_AddService(c *fiber.Ctx) error {
	var (
		service database.Service
		db      *sql.DB
		found   bool
		err     error
	)

	db = *(c.Locals("database").(**sql.DB))

	if c.BodyParser(&service) != nil {
		return util.ErrBadJSON(c)
	}

	if service.Name == "" || service.Desc == "" || service.Url == "" {
		return util.ErrBadData(c)
	}

	if found, err = service.Get(db, service.Name); err != nil {
		util.Fail("error while searching for a service (\"%s\"): %s", service.Name, err.Error())
		return util.ErrServer(c)
	}

	if found {
		return util.ErrEntryExists(c)
	}

	if err = service.Save(db); err != nil {
		util.Fail("error while saving a new service (\"%s\"): %s", service.Name, err.Error())
		return util.ErrServer(c)
	}

	return util.NoError(c)
}

func DEL_RemovePost(c *fiber.Ctx) error {
	var (
		db    *sql.DB
		id    string
		found bool
		err   error
		post  database.Post
	)

	db = *(c.Locals("database").(**sql.DB))

	if id = c.Query("id"); id == "" {
		return util.ErrBadData(c)
	}

	if found, err = post.Get(db, id); err != nil {
		util.Fail("error while searching for a post (\"%s\"): %s", id, err.Error())
		return util.ErrServer(c)
	}

	if !found {
		return util.ErrEntryNotExists(c)
	}

	if err = post.Remove(db); err != nil {
		util.Fail("error while removing a post (\"%s\"): %s", post.ID, err.Error())
		return util.ErrServer(c)
	}

	return util.NoError(c)
}

func PUT_AddPost(c *fiber.Ctx) error {
	var (
		db   *sql.DB
		post database.Post
		err  error
	)

	db = *(c.Locals("database").(**sql.DB))
	post.Public = 1

	if c.BodyParser(&post) != nil {
		return util.ErrBadJSON(c)
	}

	if post.Title == "" || post.Author == "" || post.Content == "" {
		return util.ErrBadData(c)
	}

	post.Date = time.Now().Format("02/01/06")

	if err = post.Save(db); err != nil && strings.Contains(err.Error(), sqlite3.ErrConstraintUnique.Error()) {
		return util.ErrEntryExists(c)
	}

	if err != nil {
		util.Fail("error while saving a new post (\"%s\"): %s", post.ID, err.Error())
		return util.ErrServer(c)
	}

	return util.NoError(c)
}

package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/database"
	"github.com/ngn13/website/api/util"
)

func getVoteHash(id string, ip string) string {
	return util.GetSHA512(id + "_" + ip)
}

func GET_VoteGet(c *fiber.Ctx) error {
	var (
		db    *sql.DB
		id    string
		hash  string
		vote  database.Vote
		found bool
		err   error
	)

	db = *(c.Locals("database").(**sql.DB))

	if id = c.Query("id"); id == "" {
		return util.ErrBadData(c)
	}

	hash = getVoteHash(id, util.GetIP(c))

	if found, err = vote.Get(db, hash); err != nil {
		util.Fail("error while searchig for a vote (\"%s\"): %s", hash, err.Error())
		return util.ErrServer(c)
	}

	if !found {
		return util.ErrEntryNotExists(c)
	}

	if vote.IsUpvote {
		return c.JSON(fiber.Map{
			"error":  "",
			"result": "upvote",
		})
	}

	return c.JSON(fiber.Map{
		"error":  "",
		"result": "downvote",
	})
}

func GET_VoteSet(c *fiber.Ctx) error {
	var (
		db        *sql.DB
		id        string
		is_upvote bool
		hash      string
		vote      database.Vote
		post      database.Post
		found     bool
		err       error
	)

	db = *(c.Locals("database").(**sql.DB))
	id = c.Query("id")

	if c.Query("to") == "" || id == "" {
		return util.ErrBadData(c)
	}

	if found, err = post.Get(db, id); err != nil {
		util.Fail("error while searching for a post (\"%s\"): %s", id, err.Error())
		return util.ErrServer(c)
	}

	if !found {
		return util.ErrEntryNotExists(c)
	}

	is_upvote = c.Query("to") == "upvote"
	hash = getVoteHash(id, util.GetIP(c))

	if found, err = vote.Get(db, hash); err != nil {
		util.Fail("error while searching for a vote (\"%s\"): %s", hash, err.Error())
		return util.ErrServer(c)
	}

	if found {
		if vote.IsUpvote == is_upvote {
			return util.ErrEntryExists(c)
		}

		if vote.IsUpvote && !is_upvote {
			post.Vote -= 2
		}

		if !vote.IsUpvote && is_upvote {
			post.Vote += 2
		}

		vote.IsUpvote = is_upvote

		if err = post.Update(db); err != nil {
			util.Fail("error while updating post (\"%s\"): %s", post.ID, err.Error())
			return util.ErrServer(c)
		}

		if err = vote.Update(db); err != nil {
			util.Fail("error while updating vote (\"%s\"): %s", vote.Hash, err.Error())
			return util.ErrServer(c)
		}

		return util.NoError(c)
	}

	vote.Hash = hash
	vote.IsUpvote = is_upvote

	if is_upvote {
		post.Vote++
	} else {
		post.Vote--
	}

	if err = post.Update(db); err != nil {
		util.Fail("error while updating post (\"%s\"): %s", post.ID, err.Error())
		return util.ErrServer(c)
	}

	if err = vote.Save(db); err != nil {
		util.Fail("error while updating vote (\"%s\"): %s", vote.Hash, err.Error())
		return util.ErrServer(c)
	}

	return util.NoError(c)
}

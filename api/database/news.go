package database

import (
	"database/sql"

	"github.com/ngn13/website/api/util"
)

type News struct {
	ID      string    `json:"id"`      // ID of the news
	title   string    `json:"-"`       // title of the news (string)
	Title   Multilang `json:"title"`   // title of the news
	Author  string    `json:"author"`  // author of the news
	Time    uint64    `json:"time"`    // when the new was published
	content string    `json:"-"`       // content of the news (string)
	Content Multilang `json:"content"` // content of the news
}

func (n *News) Supports(lang string) bool {
	return n.Content.Supports(lang) && n.Title.Supports(lang)
}

func (n *News) Load() (err error) {
	if err = n.Title.Load(n.title); err != nil {
		return err
	}

	if err = n.Content.Load(n.content); err != nil {
		return err
	}

	return nil
}

func (n *News) Dump() (err error) {
	if n.title, err = n.Title.Dump(); err != nil {
		return err
	}

	if n.content, err = n.Content.Dump(); err != nil {
		return err
	}

	return nil
}

func (n *News) Scan(rows *sql.Rows) (err error) {
	err = rows.Scan(
		&n.ID, &n.title, &n.Author,
		&n.Time, &n.content)

	if err != nil {
		return err
	}

	return n.Load()
}

func (n *News) IsValid() bool {
	return n.Author != "" && n.ID != "" && !n.Title.Empty() && !n.Content.Empty()
}

func (db *Type) NewsNext(n *News) bool {
	var err error

	if nil == db.rows {
		if db.rows, err = db.sql.Query("SELECT * FROM news"); err != nil {
			util.Fail("failed to query news table: %s", err.Error())
			goto fail
		}
	}

	if !db.rows.Next() {
		goto fail
	}

	if err = n.Scan(db.rows); err != nil {
		util.Fail("failed to scan the news table: %s", err.Error())
		goto fail
	}

	return true

fail:
	if db.rows != nil {
		db.rows.Close()
	}
	db.rows = nil

	return false
}

func (db *Type) NewsRemove(id string) error {
	_, err := db.sql.Exec(
		"DELETE FROM news WHERE id = ?",
		id,
	)

	return err
}

func (db *Type) NewsAdd(n *News) (err error) {
	if err = n.Dump(); err != nil {
		return err
	}

	_, err = db.sql.Exec(
		`INSERT OR REPLACE INTO news(
			id, title, author, time, content
		) values(?, ?, ?, ?, ?)`,
		n.ID, n.title,
		n.Author, n.Time, n.content,
	)

	return err
}

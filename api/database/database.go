package database

import (
	"database/sql"
	"github.com/ngn13/website/api/global"
)

type Type struct {
	Sql   *sql.DB
	Votes []global.Vote
}

func (t *Type) Setup() error {
	t.Votes = []global.Vote{}

	_, err := t.Sql.Exec(`
    CREATE TABLE IF NOT EXISTS posts(
      id      TEXT NOT NULL UNIQUE,
      title   TEXT NOT NULL,
      author  TEXT NOT NULL,
      date    TEXT NOT NULL,
      content TEXT NOT NULL,
      public  INTEGER NOT NULL,
      vote    INTEGER NOT NULL
    );
  `)

	if err != nil {
		return err
	}

	_, err = t.Sql.Exec(`
    CREATE TABLE IF NOT EXISTS services( 
      name    TEXT NOT NULL UNIQUE,
      desc    TEXT NOT NULL,
      url     TEXT NOT NULL
    );
  `)

	return err
}

func (t *Type) Open(p string) error {
	var err error

	if t.Sql, err = sql.Open("sqlite3", p); err != nil {
		return err
	}

	return t.Setup()
}

func (t *Type) Close() {
	t.Sql.Close()
}

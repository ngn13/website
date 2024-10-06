package database

import (
	"database/sql"
)

func Setup(db *sql.DB) error {
	_, err := db.Exec(`
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

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS services( 
      name    TEXT NOT NULL UNIQUE,
      desc    TEXT NOT NULL,
      url     TEXT NOT NULL
    );
  `)

	if err != nil {
		return err
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS votes( 
      hash       TEXT NOT NULL UNIQUE,
      is_upvote  INTEGER NOT NULL
    );
  `)

	return err
}

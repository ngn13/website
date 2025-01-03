package database

import (
	"fmt"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Type struct {
	sql  *sql.DB
	rows *sql.Rows
}

func (db *Type) Load() (err error) {
	if db.sql, err = sql.Open("sqlite3", "data.db"); err != nil {
		return fmt.Errorf("cannot access the database: %s", err.Error())
	}

	// see database/service.go
	_, err = db.sql.Exec(`
    CREATE TABLE IF NOT EXISTS services(
      name       TEXT NOT NULL UNIQUE,
      desc       TEXT NOT NULL,
			check_time INTEGER NOT NULL,
			check_res  INTEGER NOT NULL,
			check_url  TEXT NOT NULL,
      clear      TEXT,
      onion      TEXT,
      i2p        TEXT
    );
  `)

	if err != nil {
		return fmt.Errorf("failed to create the services table: %s", err.Error())
	}

	// see database/news.go
	_, err = db.sql.Exec(`
    CREATE TABLE IF NOT EXISTS news(
			id      TEXT NOT NULL UNIQUE,
			title   TEXT NOT NULL,
			author  TEXT NOT NULL,
			time    INTEGER NOT NULL,
			content TEXT NOT NULL
		);
  `)

	if err != nil {
		return fmt.Errorf("failed to create the news table: %s", err.Error())
	}

	// see database/admin.go
	_, err = db.sql.Exec(`
    CREATE TABLE IF NOT EXISTS admin_log(
			action TEXT NOT NULL,
			time   INTEGER NOT NULL
    );
  `)

	if err != nil {
		return fmt.Errorf("failed to create the admin_log table: %s", err.Error())
	}

	return nil
}

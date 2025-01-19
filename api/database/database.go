package database

import (
	"fmt"
	"os"
	"path"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	SQL_PATH = "sql"

	TABLE_ADMIN_LOG = "admin_log" // stores administrator logs
	TABLE_METRICS   = "metrics"   // stores API usage metrcis
	TABLE_NEWS      = "news"      // stores news posts
	TABLE_SERVICES  = "services"  // stores services
	TABLE_PROJECTS  = "projects"  // stores projects
)

var tables []string = []string{
	TABLE_ADMIN_LOG, TABLE_METRICS, TABLE_NEWS,
	TABLE_SERVICES, TABLE_PROJECTS,
}

type Type struct {
	sql  *sql.DB
	rows *sql.Rows
}

func (db *Type) create_table(table string) error {
	var (
		err   error
		query []byte
	)

	query_path := path.Join(SQL_PATH, table+".sql")

	if query, err = os.ReadFile(query_path); err != nil {
		return fmt.Errorf("failed to read %s for table %s: %", query_path, table, err.Error())
	}

	if _, err = db.sql.Exec(string(query)); err != nil {
		return fmt.Errorf("failed to create the %s table: %s", table, err.Error())
	}

	return nil
}

func (db *Type) Load() (err error) {
	if db.sql, err = sql.Open("sqlite3", "data.db"); err != nil {
		return fmt.Errorf("failed access the database: %s", err.Error())
	}

	for _, table := range tables {
		if err = db.create_table(table); err != nil {
			return err
		}
	}

	return nil
}

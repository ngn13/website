package database

import (
	"database/sql"

	"github.com/ngn13/website/api/util"
)

func (db *Type) MetricsGet(key string) (uint64, error) {
	var (
		row   *sql.Row
		count uint64
		err   error
	)

	if row = db.sql.QueryRow("SELECT value FROM "+TABLE_METRICS+" WHERE key = ?", key); row == nil {
		return 0, nil
	}

	if err = row.Scan(&count); err != nil && err != sql.ErrNoRows {
		util.Fail("failed to scan the table: %s", err.Error())
		return 0, err
	}

	if err == sql.ErrNoRows {
		return 0, nil
	}

	return count, nil
}

func (db *Type) MetricsSet(key string, value uint64) error {
	var (
		err error
		res sql.Result
	)

	if res, err = db.sql.Exec("UPDATE "+TABLE_METRICS+" SET value = ? WHERE key = ?", value, key); err != nil && err != sql.ErrNoRows {
		util.Fail("failed to query table: %s", err.Error())
		return err
	}

	if effected, err := res.RowsAffected(); err != nil {
		return err
	} else if effected < 1 {
		_, err = db.sql.Exec(
			`INSERT INTO "+TABLE_METRICS+"(
			  key, value
		  ) values(?, ?)`,
			key, value,
		)

		return err
	}

	return nil
}

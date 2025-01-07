package database

import (
	"database/sql"
)

func (db *Type) VisitorGet() (uint64, error) {
	var (
		row   *sql.Row
		count uint64
		err   error
	)

	if row = db.sql.QueryRow("SELECT count FROM visitor_count WHERE id = 0"); row == nil {
		return 0, nil
	}

	if err = row.Scan(&count); err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if err == sql.ErrNoRows {
		return 0, nil
	}

	return count, nil
}

func (db *Type) VisitorIncrement() (err error) {
	if _, err = db.sql.Exec("UPDATE visitor_count SET count = count + 1 WHERE id = 0"); err != nil && err != sql.ErrNoRows {
		return err
	}

	// TODO: err is always nil even if there is no rows for some reason, check sql.Result instead

	if err == sql.ErrNoRows {
		_, err = db.sql.Exec(
			`INSERT INTO visitor_count(
			  id, count
		  ) values(?, ?)`,
			0, 0,
		)

		return err
	}

	return nil
}

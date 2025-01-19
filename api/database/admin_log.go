package database

import (
	"database/sql"
	"fmt"

	"github.com/ngn13/website/api/util"
)

type AdminLog struct {
	Action string `json:"action"` // action that was performed (service removal, service addition etc.)
	Time   int64  `json:"time"`   // time when the action was performed
}

func (l *AdminLog) Scan(rows *sql.Rows) (err error) {
	if rows != nil {
		return rows.Scan(&l.Action, &l.Time)
	}

	return fmt.Errorf("no row/rows specified")
}

func (db *Type) AdminLogNext(l *AdminLog) bool {
	var err error

	if nil == db.rows {
		if db.rows, err = db.sql.Query("SELECT * FROM " + TABLE_ADMIN_LOG); err != nil {
			util.Fail("failed to query table: %s", err.Error())
			goto fail
		}
	}

	if !db.rows.Next() {
		goto fail
	}

	if err = l.Scan(db.rows); err != nil {
		util.Fail("failed to scan the table: %s", err.Error())
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

func (db *Type) AdminLogAdd(l *AdminLog) error {
	_, err := db.sql.Exec(
		"INSERT INTO "+TABLE_ADMIN_LOG+`(
			action, time
		) values(?, ?)`,
		&l.Action, &l.Time,
	)

	return err
}

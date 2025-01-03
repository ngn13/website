package database

import (
	"database/sql"
	"fmt"

	"github.com/ngn13/website/api/util"
)

type Service struct {
	Name      string    `json:"name"`       // name of the service
	desc      string    `json:"-"`          // description of the service (string)
	Desc      Multilang `json:"desc"`       // description of the service
	CheckTime uint64    `json:"check_time"` // last status check time
	CheckRes  uint8     `json:"check_res"`  // result of the status check
	CheckURL  string    `json:"check_url"`  // URL used for status check
	Clear     string    `json:"clear"`      // Clearnet (cringe) URL for the service
	Onion     string    `json:"onion"`      // Onion (TOR) URL for the service
	I2P       string    `json:"i2p"`        // I2P URL for the service
}

func (s *Service) Load() error {
	return s.Desc.Load(s.desc)
}

func (s *Service) Dump() (err error) {
	s.desc, err = s.Desc.Dump()
	return
}

func (s *Service) Scan(rows *sql.Rows, row *sql.Row) (err error) {
	if rows != nil {
		err = rows.Scan(
			&s.Name, &s.desc,
			&s.CheckTime, &s.CheckRes, &s.CheckURL,
			&s.Clear, &s.Onion, &s.I2P)
	} else if row != nil {
		err = row.Scan(
			&s.Name, &s.desc,
			&s.CheckTime, &s.CheckRes, &s.CheckURL,
			&s.Clear, &s.Onion, &s.I2P)
	} else {
		return fmt.Errorf("no row/rows specified")
	}

	if err != nil {
		return err
	}

	return s.Load()
}

func (s *Service) IsValid() bool {
	return s.Name != "" && (s.Clear != "" || s.Onion != "" || s.I2P != "") && !s.Desc.Empty()
}

func (db *Type) ServiceNext(s *Service) bool {
	var err error

	if nil == db.rows {
		if db.rows, err = db.sql.Query("SELECT * FROM services"); err != nil {
			util.Fail("failed to query services table: %s", err.Error())
			goto fail
		}
	}

	if !db.rows.Next() {
		goto fail
	}

	if err = s.Scan(db.rows, nil); err != nil {
		util.Fail("failed to scan the services table: %s", err.Error())
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

func (db *Type) ServiceFind(name string) (*Service, error) {
	var (
		row *sql.Row
		s   Service
		err error
	)

	if row = db.sql.QueryRow("SELECT * FROM services WHERE name = ?", name); row == nil || row.Err() == sql.ErrNoRows {
		return nil, nil
	}

	if err = s.Scan(nil, row); err != nil {
		return nil, err
	}

	return &s, nil
}

func (db *Type) ServiceRemove(name string) error {
	_, err := db.sql.Exec(
		"DELETE FROM services WHERE name = ?",
		name,
	)

	return err
}

func (db *Type) ServiceUpdate(s *Service) (err error) {
	if err = s.Dump(); err != nil {
		return err
	}

	_, err = db.sql.Exec(
		`INSERT OR REPLACE INTO services(
			name, desc, check_time, check_res, check_url, clear, onion, i2p
		) values(?, ?, ?, ?, ?, ?, ?, ?)`,
		s.Name, s.desc,
		s.CheckTime, s.CheckRes, s.CheckURL,
		s.Clear, s.Onion, s.I2P,
	)

	return err
}

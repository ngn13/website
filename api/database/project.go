package database

import (
	"database/sql"

	"github.com/ngn13/website/api/util"
)

type Project struct {
	Name    string    `json:"name"`    // name of the project
	desc    string    `json:"-"`       // description of the project (string)
	Desc    Multilang `json:"desc"`    // description of the project
	URL     string    `json:"url"`     // URL of the project's homepage/source
	License string    `json:"license"` // name of project's license
}

func (p *Project) Load() error {
	return p.Desc.Load(p.desc)
}

func (p *Project) Dump() (err error) {
	p.desc, err = p.Desc.Dump()
	return
}

func (p *Project) Scan(rows *sql.Rows) (err error) {
	if err = rows.Scan(
		&p.Name, &p.desc,
		&p.URL, &p.License); err != nil {
		return err
	}

	return p.Load()
}

func (p *Project) IsValid() bool {
	return p.Name != "" && p.URL != "" && !p.Desc.Empty()
}

func (db *Type) ProjectNext(p *Project) bool {
	var err error

	if nil == db.rows {
		if db.rows, err = db.sql.Query("SELECT * FROM " + TABLE_PROJECTS); err != nil {
			util.Fail("failed to query table: %s", err.Error())
			goto fail
		}
	}

	if !db.rows.Next() {
		goto fail
	}

	if err = p.Scan(db.rows); err != nil {
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

func (db *Type) ProjectRemove(name string) error {
	_, err := db.sql.Exec(
		"DELETE FROM "+TABLE_PROJECTS+" WHERE name = ?",
		name,
	)

	return err
}

func (db *Type) ProjectAdd(p *Project) (err error) {
	if err = p.Dump(); err != nil {
		return err
	}

	_, err = db.sql.Exec(
		"INSERT OR REPLACE INTO "+TABLE_PROJECTS+`(
      name, desc, url, license
		) values(?, ?, ?, ?)`,
		p.Name, p.desc, p.URL, p.License,
	)

	return err
}

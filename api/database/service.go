package database

import (
	"database/sql"
)

type Service struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Url  string `json:"url"`
}

func (s *Service) Load(rows *sql.Rows) error {
	return rows.Scan(&s.Name, &s.Desc, &s.Url)
}

func (s *Service) Get(db *sql.DB, name string) (bool, error) {
	var (
		success bool
		rows    *sql.Rows
		err     error
	)

	if rows, err = db.Query("SELECT * FROM services WHERE name = ?", name); err != nil {
		return false, err
	}
	defer rows.Close()

	if success = rows.Next(); !success {
		return false, nil
	}

	if err = s.Load(rows); err != nil {
		return false, err
	}

	return true, nil
}

func (s *Service) Remove(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM services WHERE name = ?",
		s.Name,
	)

	return err
}

func (s *Service) Save(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO services(name, desc, url) values(?, ?, ?)",
		s.Name, s.Desc, s.Url,
	)

	return err
}

package database

import "database/sql"

type Vote struct {
	Hash     string
	IsUpvote bool
}

func (v *Vote) Load(rows *sql.Rows) error {
	return rows.Scan(&v.Hash, &v.IsUpvote)
}

func (v *Vote) Get(db *sql.DB, hash string) (bool, error) {
	var (
		success bool
		rows    *sql.Rows
		err     error
	)

	if rows, err = db.Query("SELECT * FROM votes WHERE hash = ?", hash); err != nil {
		return false, err
	}
	defer rows.Close()

	if success = rows.Next(); !success {
		return false, nil
	}

	if err = v.Load(rows); err != nil {
		return false, err
	}

	return true, nil
}

func (v *Vote) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE votes SET is_upvote = ? WHERE hash = ?", v.IsUpvote, v.Hash)
	return err
}

func (v *Vote) Save(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO votes(hash, is_upvote) values(?, ?)",
		v.Hash, v.IsUpvote,
	)

	return err
}

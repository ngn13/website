package database

import (
	"database/sql"

	"github.com/ngn13/website/api/util"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Date    string `json:"date"`
	Content string `json:"content"`
	Public  int    `json:"public"`
	Vote    int    `json:"vote"`
}

func (p *Post) Load(rows *sql.Rows) error {
	return rows.Scan(&p.ID, &p.Title, &p.Author, &p.Date, &p.Content, &p.Public, &p.Vote)
}

func (p *Post) Get(db *sql.DB, id string) (bool, error) {
	var (
		success bool
		rows    *sql.Rows
		err     error
	)

	if rows, err = db.Query("SELECT * FROM posts WHERE id = ?", id); err != nil {
		return false, err
	}
	defer rows.Close()

	if success = rows.Next(); !success {
		return false, nil
	}

	if err = p.Load(rows); err != nil {
		return false, err
	}

	return true, nil
}

func (p *Post) Remove(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", p.ID)
	return err
}

func (p *Post) Save(db *sql.DB) error {
	p.ID = util.TitleToID(p.Title)

	_, err := db.Exec(
		"INSERT INTO posts(id, title, author, date, content, public, vote) values(?, ?, ?, ?, ?, ?, ?)",
		p.ID, p.Title, p.Author, p.Date, p.Content, p.Public, p.Vote,
	)

	return err
}

func (p *Post) Update(db *sql.DB) error {
	p.ID = util.TitleToID(p.Title)

	_, err := db.Exec(
		"UPDATE posts SET title = ?, author = ?, date = ?, content = ?, public = ?, vote = ? WHERE id = ?",
		p.Title, p.Author, p.Date, p.Content, p.Public, p.Vote, p.ID,
	)

	return err
}

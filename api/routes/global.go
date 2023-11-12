package routes

import (
	"database/sql"
	"strings"
)

// ############### BLOG ###############
type Post struct {
  ID      string    `json:"id"`
  Title   string    `json:"title"` 
  Author  string    `json:"author"` 
  Date    string    `json:"date"`
  Content string    `json:"content"`
  Public  int       `json:"public"`
  Vote    int       `json:"vote"`
}

var votelist = []Vote{}
type Vote struct {
  Post    string
  Client  string 
  Status  string
}

func PostFromRow(post *Post, rows *sql.Rows) error{
  err := rows.Scan(&post.ID, &post.Title, &post.Author, &post.Date, &post.Content, &post.Public, &post.Vote)
  if err != nil {
    return err
  }

  return nil
}

func GetPostByID(id string) (Post, string) {
  var post Post = Post{}
  post.Title = "NONE"

  rows, err := DB.Query("SELECT * FROM posts WHERE id = ?", id)

  if err != nil{
    return post, "Server error" 
  }

  success := rows.Next()
  if !success {
    rows.Close()
    return post, "Post not found"
  }

  err = PostFromRow(&post, rows)
  if err != nil {
    rows.Close()
    return post, "Server error" 
  }
  rows.Close()

  if post.Title == "NONE" {
    return post, "Post not found" 
  }

  return post, "" 
}

func TitleToID(name string) string{
  return strings.ToLower(strings.ReplaceAll(name, " ", ""))
}

// ############### SERVICES ############### 
type Service struct {
  Name    string    `json:"name"`
  Desc    string    `json:"desc"`
  Url     string    `json:"url"`
}

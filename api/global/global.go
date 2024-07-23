package global

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Date    string `json:"date"`
	Content string `json:"content"`
	Public  int    `json:"public"`
	Vote    int    `json:"vote"`
}

type Service struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Url  string `json:"url"`
}

type Vote struct {
	Post   string
	Client string
	Status string
}

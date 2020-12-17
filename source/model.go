package db

type Article struct {
	Id int64 `json:"id,omitempty"`

	Title string `json:"title,omitempty"`

	Username string `json:"username,omitempty"`

	Tags []Tag `json:"tags,omitempty"`

	Date string `json:"date,omitempty"`

	Content string `json:"content,omitempty"`

	Comments []Comment `json:"comments,omitempty"`
}


type Tag struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

type User struct {
	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`
}

type Comment struct {
	User string `json:"user,omitempty"`

	ArticleId int64 `json:"article_id,omitempty"`

	Date string `json:"date,omitempty"`

	Content string `json:"content,omitempty"`
}
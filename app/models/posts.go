package models

type Posts struct {
	ArticleID uint64 `db:"article_uuid"`
	Title     string `db:"title"`
	Content   string `db:"content"`
	Author    string `db:"author"`
}

type JsonPost struct {
	ArticleID uint64 `json:"article_uuid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
}

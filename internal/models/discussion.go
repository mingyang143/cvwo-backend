package models

type Discussion struct {
	ID   int64    `json:"id"`
	UserId int64 	`json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Likes int64 `json:"likes"`
	Comments []string `json:"comments"`

}


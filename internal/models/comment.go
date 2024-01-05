package models

type Comment struct {
	ID   int64    `json:"id"`
	Comment string 	`json:"comment"`
	DiscussionId int64 `json:"discussionId"`
}


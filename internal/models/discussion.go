package models

type Discussion struct {
	ID   int64    `json:"id"`
	UserId int64 	`json:"userId"`
	Title string `json:"title"`
	Content string `json:"content"`
	Likes int64 `json:"likes"`
	Comments []Comment `json:"comments"`
	Tags []Tag `json:"tags"`

}



type DiscussionId struct {
	DiscussionId int64 `json:"discussionId"`
}
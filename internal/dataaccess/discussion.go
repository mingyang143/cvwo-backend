package dataaccess

import (
	"fmt"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)


func EditDiscussion(db *database.Database, discussion models.Discussion) (models.Discussion, error) {
	fmt.Printf("call editdiscussion(db, discussion)\n")
	return db.EditDiscussion(discussion)
}

func ListDiscussion(db *database.Database) ([]models.Discussion, error) {
	fmt.Printf("call AllDiscussion()\n")
	discussion := db.AllDiscussion()
	
	return discussion, nil
}
func CreateDiscussion(db *database.Database, discussion models.Discussion) (models.Discussion, error){
	fmt.Printf("call CreateDiscussion()\n")
	
	
	return db.AddDiscussion(discussion)
}

func AddComment(db *database.Database, comment models.Comment) (models.Comment, error) {
	fmt.Printf("call add(db, discussion)\n")
	return db.AddComment(comment)
}

func IncreaseLikes(db *database.Database, discussionId int64) error {
	return db.IncreaseLikes(discussionId)
}

// func AddDiscussion(db *database.Database, discussion models.Discussion) (models.Discussion, error) {
// 	fmt.Printf("call add(db, discussion)\n")
// 	return db.AddDiscussion(discussion)
// }
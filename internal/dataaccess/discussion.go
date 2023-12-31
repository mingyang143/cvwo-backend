package dataaccess

import (
	"fmt"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)


func AddDiscussion(db *database.Database, discussion models.Discussion) (models.Discussion, error) {
	fmt.Printf("call add(db, discussion)\n")
	return db.AddDiscussion(discussion)
}

func EditDiscussion(db *database.Database, discussion models.Discussion) (models.Discussion, error) {
	fmt.Printf("call editdiscussion(db, discussion)\n")
	return db.EditDiscussion(discussion)
}

func ListDiscussion(db *database.Database) ([]models.Discussion, error) {
	fmt.Printf("call AllDiscussion()\n")
	discussion := db.AllDiscussion()
	
	return discussion, nil
}

// func AddComment(db *database.Database, discussion models.Discussion) (models.Discussion, error) {
// 	fmt.Printf("call add(db, discussion)\n")
// 	return db.AddComment(discussion)
// }
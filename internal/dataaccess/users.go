package users

import (
	"fmt"

	"github.com/CVWO/sample-go-app/internal/database"
	"github.com/CVWO/sample-go-app/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	// users := []models.User{
	// 	{
	// 		ID:   1,
	// 		Name: "CVWO",
	// 	},
	// }
	fmt.Printf("call AllUsers()\n")
	users := db.AllUser()
	
	return users, nil
}

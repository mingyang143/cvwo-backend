package database

import (
	"database/sql"
	"fmt"

	"github.com/CVWO/sample-go-app/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB;
}

func GetDB() (*Database, error) {
	name := "mysql"
	connectStr := "user1:password@tcp(localhost:3306)/internetforum"

	db, err :=sql.Open(name, connectStr)
	
	if err != nil {
		return nil, err
	}
	database := Database{
		db: db,
	}

	return &database, nil
}

func (db *Database) Close() {
	db.Close()
}


func (db *Database) AllUser() []models.User {
	rows, err := db.db.Query("select id, username from users")
	if err != nil {
		fmt.Printf("failed to get the list of users %s\n", err)
	}
	defer rows.Close()

	var users []models.User
	
	var (
		id int
		name string
	)
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Printf("failed to retrieve values for user %s \n", err)
		}
		u := models.User{
			ID: id,
			Name: name,
		}
		users = append(users, u)
	}
	return users
}
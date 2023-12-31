package database

import (
	"context"
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
		id int64
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



func (db *Database) AddUser(user models.User) (models.User, error) {
	insertResult, err := db.db.ExecContext(context.Background(),"insert into users (username) values (?)", user.Name)
	if err != nil {
		fmt.Printf("failed to add users %s\n", err)
		return models.User{}, err
	}
	id, err := insertResult.LastInsertId()
	return models.User{
		ID: id,
		Name: user.Name,

	}, nil
}

func (db *Database) IsValidUser(user models.User) (models.User,error) {
	var id int64
	if err := db.db.QueryRow("select id from users where username=?", user.Name).Scan(&id); err !=nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found %s", user.Name)
		} 
		return models.User{}, fmt.Errorf("user is invalid %s", user.Name)
	}
	return models.User{
		ID: id,
		Name: user.Name,
	}, nil
}

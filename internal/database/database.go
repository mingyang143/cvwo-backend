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


func (db *Database) AllDiscussion() []models.Discussion {
	rows, err := db.db.Query("select id,user_id,title,content,likes from discussion")
	if err != nil {
		fmt.Printf("failed to get discussion %s\n", err)
	}
	defer rows.Close()

	var discussions []models.Discussion
	var commentAll []string
	var (
		id int64
		user_id int64
		title string
		content string
		likes int64
		
	)
	for rows.Next() {
		
		err := rows.Scan(&id,&user_id,&title, &content,&likes)
		if err != nil {
			fmt.Printf("failed to retrieve discussion %s \n", err)
		}
		
		rowComments, errComments := db.db.Query("select comment from comments where discussion_id=?", id)
		if errComments != nil {
		fmt.Printf("failed to get comments %s\n", err)
		}
		defer rowComments.Close()
		var(
			comment string
		)
		for rowComments.Next(){
			err := rowComments.Scan(&comment)
			if err != nil {
				fmt.Printf("failed to retrieve comment %s \n", err)
			}
		commentAll = append(commentAll, comment)
		}


		p := models.Discussion{
			ID: id,
			UserId: user_id,
			Title: title,
			Content: content,
			Likes: likes,
			Comments: commentAll,
			

		}
		discussions = append(discussions, p)
	}
	return discussions
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


func (db *Database) EditDiscussion(discussion models.Discussion) (models.Discussion, error) {
	
	updateResult, err := db.db.ExecContext(context.Background(), "update discussion set title=?, content=? where id=?", 
	discussion.Title, discussion.Content, discussion.ID)

	if err != nil {
		fmt.Printf("Failed to add discussion %s\n", err)
	}

	n, err := updateResult.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to add discussion %s\n", err)
	}
	fmt.Printf("%d discussion updated\n", n);

	return discussion, nil
}


// func (db *Database) AddComment(discussion models.Discussion) (models.Discussion, error) {
// 	insertResult, err := db.db.ExecContext(context.Background(),"insert into comments(comment, discussion_id) values(?)", discussion.ID, comment )
// 	if err != nil {
// 		fmt.Printf("failed to add comment %s\n", err)
// 		return models.Discussion{}, err
// 	}
	
// 	id, err := insertResult.LastInsertId()
// 	return models.Discussion{
		
// 	}, nil
// }


// func (db *Database) AddDiscussion(discussion models.Discussion) (models.Discussion, error) {
// 	//todo
// 	return models.Discussion{}, nil
// }
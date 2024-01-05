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
	db.db.Close()
}

//ask dad how to change comment to a object and get the object from backend to react
func (db *Database) AllDiscussion() []models.Discussion {
	rows, err := db.db.Query("select id,user_id,title,content,likes from discussion")
	if err != nil {
		fmt.Printf("failed to get discussion %s\n", err)
	}
	defer rows.Close()

	var discussions = make([]models.Discussion,0) 
	//var commentAll []string
	var (
		id int64
		user_id int64
		title string
		content string
		likes int64
		
	)
	for rows.Next() {
		comments := make([]models.Comment, 0)
		var discussion_id int64
		var comment string
		
		err := rows.Scan(&id,&user_id,&title, &content,&likes)
		if err != nil {
			fmt.Printf("failed to retrieve discussion %s \n", err)
		}
		
		rowComments, errComments := db.db.Query("select * from comments where discussion_id=?", id)
		if errComments != nil {
		fmt.Printf("failed to get comments %s\n", err)
		}
		defer rowComments.Close()

		for rowComments.Next(){
			var commentId int64
			err := rowComments.Scan(&commentId,&comment,&discussion_id)
			if err != nil {
				fmt.Printf("failed to retrieve comment %s \n", err)
			}
			c := models.Comment{
				ID: commentId,
				Comment: comment,
				DiscussionId: discussion_id,
			}
			comments = append(comments, c)
		}

		p := models.Discussion{
			ID: id,
			UserId: user_id,
			Title: title,
			Content: content,
			Likes: likes,
			Comments: comments,
			

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

	if err != nil {
		fmt.Printf("failed to insert Users table %s\n", err)
		return models.User{}, err
	}
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
	fmt.Printf("discussion Id to be updated %d\n", discussion.ID)
	
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

	return discussion,nil
}


func (db *Database) AddDiscussion(discussion models.Discussion) (models.Discussion, error) {
	fmt.Printf("input for discussion %v\n", discussion)
	fmt.Printf("userId for discussion %d\n", discussion.UserId)

	insertResult, err := db.db.ExecContext(context.Background(),"insert into discussion (User_id, Title, Content, Likes) values (?,?,?,?)", discussion.UserId, discussion.Title, discussion.Content, discussion.Likes)
	if err != nil {
		fmt.Printf("failed to add discussion %s\n", err)
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		fmt.Printf("failed to insert Discussion table %s\n", err)
		return models.Discussion{}, err
	}
	
	return models.Discussion{
		ID: id,
		UserId: discussion.UserId,
		Title: discussion.Title,
		Content: discussion.Content,
		Comments: make([]models.Comment,0),
	}, nil
}



func (db *Database) AddComment(comment models.Comment) (models.Comment, error) {
	insertResult, err := db.db.ExecContext(context.Background(),"insert into comments(comment, discussion_id) values(?,?)", comment.Comment, comment.DiscussionId )
	if err != nil {
		fmt.Printf("failed to add comment %s\n", err)
		return models.Comment{}, err
	}
	
	id, err := insertResult.LastInsertId()
	if err != nil {
		fmt.Printf("failed to insert comments table %s\n", err)
		return models.Comment{}, err
	}

	return models.Comment{
		ID: id,
		Comment: comment.Comment,
		DiscussionId: comment.DiscussionId,
	}, nil
}

func (db *Database) IncreaseLikes(discussionId int64) error {
	var likes int64
	if err := db.db.QueryRow("select likes from discussion where id=?", discussionId).Scan(&likes); err !=nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no dicussion found for id:%ds", discussionId)
		} 
		return fmt.Errorf("discussion id is invalid %d", discussionId)
	}

	_, err :=db.db.ExecContext(context.Background(), "update discussion set likes=? where id=?", likes + 1, discussionId)
	if err != nil {
		return err
	}
	return nil
}

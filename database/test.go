package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user1:Pas$w0rd@tcp(localhost:3306)/internetforum")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    insert, err := db.Query("insert into discussion (user, title, content) values('mingyang','12345','12345')")
    if err !=nil {
        panic(err.Error())
    }
    defer insert.Close()
    fmt.Println("Yay, values added!")

}
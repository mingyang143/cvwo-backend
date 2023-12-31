package models

import "fmt"

type User struct {
	ID   int64    `json:"id"`
	Name string `json:"name"`
}

func (user *User) Greet() string {
	return fmt.Sprintf("Hello, I am %s", user.Name)
}



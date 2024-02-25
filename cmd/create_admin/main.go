package main

import (
	"context"
	"fmt"
	"gruffalinas/internal/services"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	c := services.NewContainer()

	if len(os.Args) < 3 {
		fmt.Println(
			"Not enough arguments - add email & password (create admin {username} {password})",
		)

		return
	}

	username := os.Args[1]
	hashed, err := bcrypt.GenerateFromPassword([]byte(os.Args[2]), bcrypt.DefaultCost)

	if err != nil {
		fmt.Printf("Error - %v", err)
	}

	sql := `INSERT INTO users
        (created_at, updated_at, username, password)
        VALUES
        (NOW(), NOW(), $1, $2)`

	if _, err := c.Db.Exec(context.Background(), sql, username, hashed); err == nil {
		fmt.Println("User ", username, " added")
	} else {
		fmt.Printf("Failed adding user: %v\n", err)
	}
}

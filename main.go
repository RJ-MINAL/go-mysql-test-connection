package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// User struct for table user testdb
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	log.Println("Go MySQL Test")

	db, err := sql.Open("mysql", "testUser:password@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully connected to MySQL Database")

	/* 	insert, err := db.Query("INSERT INTO user(name) VALUES('Rolando')")

	   	if err != nil {
	   		panic(err.Error())
	   	}

	   	defer insert.Close()

	   	fmt.Println("Successfully inserted into user table") */

	// multiple row query
	results, err := db.Query("SELECT id, name from user")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.ID, user.Name)
	}

	// Execute single row query
	var tag User
	err = db.QueryRow("SELECT id, name FROM user where id = ?", 1).Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println(tag.ID)
	fmt.Println(tag.Name)
}

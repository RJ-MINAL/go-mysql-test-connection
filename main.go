package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// User struct for table user testdb
type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySQL Test")

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

	results, err := db.Query("SELECT name from user")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}
}

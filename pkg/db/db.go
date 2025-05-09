package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	fmt.Println("Database connection...")
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		fmt.Println(err)
		panic("could not connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateUserTable()
	CreateEventTable()
	CreateRegistrationTable()
}

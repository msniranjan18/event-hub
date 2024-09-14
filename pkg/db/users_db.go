package db

import "fmt"

func CreateUserTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		pwd TEXT NOT NULL
	)
	`
	_, err := DB.Exec(query)
	if err != nil {
		fmt.Println(err)
		panic("could not create users table")
	}
}


package db

import "fmt"

func CreateEventTable() {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTERGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	_, err := DB.Exec(query)
	if err != nil {
		fmt.Println(err)
		panic("could not create events table")
	}
}

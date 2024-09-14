package db

import "fmt"

func CreateRegistrationTable() {
	query := `
		CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTERGER NOT NULL,
		user_id INTERGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
		FOREIGN KEY (event_id) REFERENCES events(id)

	)
	`
	_, err := DB.Exec(query)
	if err != nil {
		fmt.Println(err)
		panic("could not create registrations table")
	}
}

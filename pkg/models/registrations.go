package models

import (
	"fmt"

	"msn.com/event-hub/pkg/db"
)

type Registration struct {
	ID      int64
	UserId  int64
	EventId int64
}

func GetAllRegistrations() ([]Registration, error) {
	var rl []Registration
	query := "SELECT * FROM registrations"
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("Error", err)
		return rl, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Registration
		err = rows.Scan(&r.ID, &r.EventId, &r.UserId)
		if err != nil {
			fmt.Println("Error", err)
			continue
		}
		rl = append(rl, r)
	}
	return rl, nil
}

func GetRegistrationsByUserId(userId int64) ([]Registration, error) {
	var rl []Registration
	query := "SELECT * FROM registrations WHERE user_id = ?"
	rows, err := db.DB.Query(query, userId)
	if err != nil {
		fmt.Println("Error", err)
		return rl, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Registration
		err = rows.Scan(&r.ID, &r.EventId, &r.UserId)
		if err != nil {
			fmt.Println("Error", err)
			continue
		}
		rl = append(rl, r)
	}
	return rl, nil
}

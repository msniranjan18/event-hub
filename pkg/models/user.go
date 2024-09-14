package models

import (
	"errors"

	"msn.com/event-hub/pkg/db"
	"msn.com/event-hub/pkg/utils"
)

type User struct {
	UserID int64
	Email  string `binding:"required"`
	PWD    string `binding:"required"`
}

func (u User) Save() (int64, error) {

	query := `
	INSERT INTO users(email, pwd)
	VALUES (?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	u.PWD, err = utils.HashPassword(u.PWD)
	if err != nil {
		return -1, err
	}

	result, err := stmt.Exec(u.Email, u.PWD)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	u.UserID = id
	return u.UserID, err
}

func (u User) Update() error {
	query := `
	UPDATE users
	SET pwd = ?
	WHERE email = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	u.PWD, _ = utils.HashPassword(u.PWD)

	_, err = stmt.Exec(u.PWD, u.Email)
	if err != nil {
		return err
	}

	return nil
}

func (u User) Delete() error {
	query := "DELETE FROM users WHERE email = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Email)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.UserID, &u.Email, &u.PWD)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}
	return users, nil
}

func GetUserByEmailId(email string) (User, error) {
	var u User
	query := "SELECT * FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, email)

	err := row.Scan(&u.UserID, &u.Email, &u.PWD)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (u User) ValidateCredentials() (int64, error) {
	query := "SELECT id, pwd FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)
	var hashedPWD string
	err := row.Scan(&u.UserID, &hashedPWD)
	if err != nil {
		return 0, errors.New("error, no such user exist")
	}

	if !utils.IsCorrectPassword(u.PWD, hashedPWD) {
		return 0, errors.New("incorrect PWD")
	}

	return u.UserID, nil
}

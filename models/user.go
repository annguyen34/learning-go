package models

import (
	"errors"
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding: "required" `
	Password string `binding: "required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		panic(err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	u.ID = userID
	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT id,password FROM users WHERE email = $1`
	row := db.DB.QueryRow(query, u.Email)
	var hashedPassword string
	err := row.Scan(&u.ID, &hashedPassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, hashedPassword)
	if !passwordIsValid {
		return errors.New("Invalid credentials")
	}
	return nil
}

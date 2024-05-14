package models

import "rest-api/db"

type User struct {
	ID       int64
	Email    string `binding: "required" `
	Password string `binding: "required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Email, u.Password)
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

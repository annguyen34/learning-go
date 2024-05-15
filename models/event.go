package models

import (
	"fmt"
	"rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserID      int64     `binding: "required"`
}

var events = []Event{}

func GetAllEvent() ([]Event, error) {
	result, err := db.DB.Query("SELECT * FROM events")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	defer result.Close()

	var events []Event

	for result.Next() {
		var e Event
		err := result.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			panic(err)
		}
		events = append(events, e)
	}
	return events, error(nil)
}

func Save(e Event) {
	query := `INSERT INTO events (name, description, location, datetime, user_id) VALUES ($1, $2, $3, $4, $5)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	fmt.Println(e)
	if err != nil {
		panic(err)
	}
	_, err = result.LastInsertId()
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = $1`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	result := stmt.QueryRow(id)
	var e Event
	err = result.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		panic(err)
	}
	return &e, err

}

func (e Event) UpdateEvent() error {
	query := `UPDATE events 
	SET name = $1, description = $2, location = $3, datetime = $4, user_id = $5
	WHERE id = $6`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	return err
}

func DeleteEvent(id int64) error {
	query := `DELETE FROM events WHERE id = $1`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

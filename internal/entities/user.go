package entities

import "time"

type User struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Birthdate  time.Time `json:"birthdate"`
	Biography  string    `json:"biography"`
	City       string    `json:"city"`
	Password   string    `db:"-"`
}

package entities

import "time"

type Post struct {
	ID     int       `json:"post_id" db:"id"`
	UserID int       `json:"user_id" db:"user_id"`
	Text   string    `json:"text" db:"text"`
	Date   time.Time `json:"date" db:"date"`
}

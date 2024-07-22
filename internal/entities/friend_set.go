package entities

import "time"

type FriendSet struct {
	IdUser1 int       `json:"id_user_1"`
	IdUser2 int       `json:"id_user_2"`
	DateAdd time.Time `json:"date_add"`
}

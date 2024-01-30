package entities

import "errors"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func (u *User) Validate() error {
	if u.Name == "" || u.Email == "" || u.Age == 0 {
		return errors.New("invalid user structure")
	}

	return nil
}

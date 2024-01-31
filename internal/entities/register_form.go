package entities

import "time"

type Form struct {
	FirstName  string    `form:"first_name" binding:"required"`
	SecondName string    `form:"second_name" binding:"required"`
	Birthdate  time.Time `form:"birthdate" time_format:"2006-01-02"`
	Biography  string    `form:"biography"`
	City       string    `form:"city"`
	Password   string    `form:"password" binding:"required"`
}

func (u *Form) Validate() error {
	// if u.Name == "" || u.Email == "" || u.Age == 0 {
	// 	return errors.New("invalid user structure")
	// }

	return nil
}

package entities

type SearchForm struct {
	FirstName  string `form:"first_name" binding:"required"`
	SecondName string `form:"second_name" binding:"required"`
}

func (u *SearchForm) Validate() error {
	// if u.Name == "" || u.Email == "" || u.Age == 0 {
	// 	return errors.New("invalid user structure")
	// }

	return nil
}

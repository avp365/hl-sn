package entities

type LoginForm struct {
	ID       int    `form:"id" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (u *LoginForm) Validate() error {
	// if u.Name == "" || u.Email == "" || u.Age == 0 {
	// 	return errors.New("invalid user structure")
	// }

	return nil
}

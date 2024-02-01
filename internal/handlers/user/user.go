package user

import (
	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/repositories"
)

func RegisterUserHandler(form *entities.Form) (int, error) {
	var user entities.User

	user.FirstName = form.FirstName
	user.SecondName = form.SecondName
	user.Biography = form.Biography
	user.Birthdate = form.Birthdate
	user.City = form.City
	user.Password = form.Password

	return repositories.UsrRep.CreateUser(user)
}

func UserGetByIdHandler(userId int) (entities.User, error) {

	return repositories.UsrRep.UserGetById(userId)

}

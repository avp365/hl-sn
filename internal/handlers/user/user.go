package user

import (
	"time"

	"github.com/avp365/hl-sn/internal/entities"
	"github.com/avp365/hl-sn/internal/repositories"
)

func RegisterUserHandler(form *entities.RegisterForm) (int, error) {
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

func UserSearchHandler(searchForm entities.SearchForm) ([]entities.User, error) {
	return repositories.UsrRep.UserSearch(searchForm.FirstName, searchForm.SecondName)
}

func FriendSet(userId, frienduserId int) (int, error) {
	return repositories.UsrRep.FriendSet(entities.FriendSet{IdUser1: userId, IdUser2: frienduserId, DateAdd: time.Now()})
}

func FriendDelete(userId, frienduserId int) (bool, error) {
	return repositories.UsrRep.FriendDelete(entities.FriendSet{IdUser1: userId, IdUser2: frienduserId, DateAdd: time.Now()})
}

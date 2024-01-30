package user

import (
	"errors"
	"fmt"
	"module09/internal/entities"
)

var idCounter = 0
var Users = make(map[int]*entities.User)
var ErrorUserNotFound = errors.New("user not found")

func Create(user entities.User) error {
	if _, ok := Users[idCounter]; ok {
		idCounter++
	}

	Users[idCounter] = &user
	Users[idCounter].ID = idCounter

	return nil
}

func GetList() []*entities.User {
	fmt.Println(Users)
	result := make([]*entities.User, 0, len(Users))
	for _, user := range Users {
		result = append(result, user)
	}

	return result
}

func GetUser(id int) (*entities.User, error) {
	user, ok := Users[id]
	if !ok {
		return nil, ErrorUserNotFound
	}

	return user, nil
}

func Delete(id int) error {
	if _, ok := Users[id]; !ok {
		return ErrorUserNotFound
	}

	delete(Users, id)
	return nil
}

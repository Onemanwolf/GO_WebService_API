package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

var (
	nextId = 3
	users  []*User
)

func GetUsers() []*User {

	return users

}

func AddUser(user User) (User, error) {

	if user.Id != 0 {

		return User{}, errors.New("User must not have an ID set")

	}

	user.Id = nextId
	nextId++

	users = append(users, &user)
	println(users)

	return user, nil
}

func GetUserById(id int) (User, error) {

	for _, user := range users {

		if user.Id == id {

			return *user, nil

		}

	}

	return User{}, fmt.Errorf("User not ID '%v' not found", id)

}

func UpdateUser(user User) (User, error) {

	for i, candidate := range users {

		if candidate.Id == user.Id {

			users[i] = &user
			return user, nil

		}
	}

	return User{}, fmt.Errorf("User not ID '%v' not found", user.Id)
}

func RemoveUserById(id int) error {

	for i, user := range users {

		if user.Id == id {

			users = append(users[:i], users[i+1:]...)
			return nil

		}
	}

	return fmt.Errorf("User not ID '%v' not found", id)
}

package userRepository

import (
	"dimi/server/mistakes"
	"errors"
)

type User struct {
	Id      int
	Name    string
	Age     int
	balance int
}

var users map[int]User

func init() {
	users = make(map[int]User, 0)
}

func GetUserById(id int) (User, error) {
	user, exists := users[id]

	if !exists {
		return User{}, mistakes.ErrNotFound
	}

	return user, nil
}

func CreateUser(user User) error {
	_, err := GetUserById(user.Id)
	if err == nil {
		return errors.New("user already exists")
	}

	return nil
}

func GetUsers() []User {
	sliceUsers := make([]User, 0, len(users))

	for _, user := range users {
		sliceUsers = append(sliceUsers, user)
	}

	return sliceUsers
}

func DeleteUser(id int) error {
	_, err := GetUserById(id)
	if err != nil {
		return err
	}
	users[id] = User{}
	return nil
}

func UpdateUser(user User) error {
	_, err := GetUserById(user.Id)

	if err != nil {
		return err
	}

	users[user.Id] = user
	return nil
}

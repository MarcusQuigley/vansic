package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	Firstname string
	Lastname  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func GetUser(id int) (User, error) {
	for _, u := range users {
		if u.Id == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with id %v not found, ya melt", id)
}

func AddUser(u User) (User, error) {
	if u.Id != 0 {
		return User{}, errors.New("User id must not have changed bawbag")
	}
	u.Id = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if u.Id == candidate.Id {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with id %v not found", u.Id)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id %v not found", id)
}

 

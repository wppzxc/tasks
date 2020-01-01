package service

import (
	"fmt"
	"github.com/wppzxc/tasks/src/database"
)

func GetUserByName(username string) (*database.Users, error) {
	user, err := database.GetUser(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByNameAndPassword(username, password string) (*database.Users, error) {
	user, err := database.GetUserByNameAndPassword(username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *database.Users) (*database.Users, error) {
	newUser, err := database.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func UpdateUser(user *database.Users) (*database.Users, error) {
	newUser, err := database.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func UpdateParentsByRegister(parent string) error {
	var users []*database.Users
	parentUser, err := GetUserByName(parent)
	if err != nil {
		return err
	}
	parentUser.SubUsers++
	users = append(users, parentUser)

	if len(parentUser.Parent) > 0 {
		TopParentUser, err := GetUserByName(parentUser.Parent)
		if err != nil {
			fmt.Println("Can't get parent for user : ", parentUser.Name)
		} else {
			TopParentUser.SubTwoUsers++
			users = append(users, TopParentUser)
		}
	}
	if err := database.UpdateUsers(users); err != nil {
		return err
	}
	return nil
}

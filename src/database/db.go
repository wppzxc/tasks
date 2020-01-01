package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Users{})
}

func GetUser(username string) (*Users, error) {
	user := &Users{}
	if err := DB.Where("name = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByNameAndPassword(username, password string) (*Users, error) {
	user := &Users{}
	if err := DB.Where("name = ? and password = ?", username, password).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *Users) (*Users, error) {
	tx := DB.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	newUser, err := GetUser(user.Name)
	if err != nil {
		return nil, fmt.Errorf("Error in get user : %s, '%s' ", user.Name, err)
	}
	return newUser, nil
}

func UpdateUser(user *Users) (*Users, error) {
	if err := DB.Model(user).Where("name = ?", user.Name).Updates(user).Error; err != nil {
		return nil, err
	}
	newUser, err := GetUser(user.Name)
	if err != nil {
		return nil, fmt.Errorf("Error in get user : %s, '%s' ", user.Name, err)
	}
	return newUser, nil
}

func UpdateUsers(users []*Users) error {
	tx := DB.Begin()
	var err error
	for _, u := range users {
		if _, err = UpdateUser(u); err != nil {
			err = fmt.Errorf("%s, Error in update user : %s", err, u.Name)
		}
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

package model

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name  string
	Email string
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("Fetched users:", users)
	return users, nil
}

func CreateUser(user *User) error {
	result := DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package repository

import (
	"fmt"
	"github.com/meetup/graph/model"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type UserRepository interface {
	FindUserByName(name string) (*User, error)
	FindUsers(limited int) ([]User, error)
	Create() error
}

type User struct {
	gorm.Model
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email"  binding:"required"`
	Meetup Meetup
}

func (user User) Create() error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user User) FindUserByName(name string) (*User, error) {
	result := db.Where("name = ?", name).First(&user)
	fmt.Println(user)
	if result.Error != nil {
		return &User{}, result.Error
	}
	return &user, nil
}

func (user User) FindUsers(count int) ([]User, error) {
	var users []User
	result := db.Limit(count).Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func ConvertUser(user *User) *model.User {
	return &model.User{
		ID:       strconv.FormatUint(uint64(user.ID), 10),
		Username: user.Name,
		Email:    user.Email,
	}
}

func ConvertUsers(users []User) []*model.User {
	var wg sync.WaitGroup
	var convertUsers []*model.User
	for _, user := range users {
		wg.Add(1)
		go func(user User) {
			convertUsers = append(convertUsers, ConvertUser(&user))
			wg.Done()
		}(user)
	}
	wg.Wait()

	return convertUsers
}

package repository

import (
	"fmt"
	"github.com/meetup/graph/model"
	"gorm.io/gorm"
	"strconv"
)

type UserEntity struct {
	gorm.Model
	Name  string `json:"name" binding:"required"`
	Email string `json:"email"  binding:"required"`
}

func (user *UserEntity) FindUserByName(name string) (*UserEntity, error) {
	result := db.Where("name = ?", name).First(&user)
	fmt.Println(user)
	if result.Error != nil {
		return &UserEntity{}, result.Error
	}
	return user, nil
}

func (user UserEntity) ConvertUser() *model.User {
	return &model.User{
		ID:       strconv.FormatUint(uint64(user.ID), 10),
		Username: user.Name,
		Email:    user.Email,
	}
}

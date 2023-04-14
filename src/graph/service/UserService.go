package service

import (
	"fmt"
	"github.com/meetup/graph/model"
	"github.com/meetup/graph/repository"
)

type UserService interface {
	GetUserByName(name string) (*model.User, error)
}

type UserServiceProvider struct{}

func (service *UserServiceProvider) GetUserByName(userName string) (*model.User, error) {
	userEntity := repository.UserEntity{}
	entity, err := userEntity.FindUserByName(userName)
	if err != nil {
		fmt.Println(err)
		return &model.User{}, err
	}
	userGraphModel := entity.ConvertUser()
	return userGraphModel, nil
}

func NewUserService() Services {
	return &UserServiceProvider{}
}

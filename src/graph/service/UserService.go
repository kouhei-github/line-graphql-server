package service

import (
	"fmt"
	"github.com/meetup/graph/model"
	"github.com/meetup/graph/repository"
)

type UserService interface {
	GetUserByName(name string) (*model.User, error)
	GetUserLimited(count int) ([]*model.User, error)
	CreateUserRecode(name string, email string) (*model.User, error)
}

type UserServiceProvider struct{}

func (service *UserServiceProvider) GetUserByName(userName string) (*model.User, error) {
	userEntity := repository.User{}
	entity, err := userEntity.FindUserByName(userName)
	if err != nil {
		fmt.Println(err)
		return &model.User{}, err
	}
	userGraphModel := repository.ConvertUser(entity)
	return userGraphModel, nil
}

func (service *UserServiceProvider) GetUserLimited(count int) ([]*model.User, error) {
	userEntity := repository.User{}
	userEntities, err := userEntity.FindUsers(count)
	if err != nil {
		fmt.Println("here")
		fmt.Println(err)
		return []*model.User{}, err
	}
	models := repository.ConvertUsers(userEntities)
	return models, nil
}

func (service *UserServiceProvider) CreateUserRecode(name string, email string) (*model.User, error) {
	userEntity := repository.User{Name: name, Email: email}
	err := userEntity.Create()
	if err != nil {
		return &model.User{}, err
	}
	return repository.ConvertUser(&userEntity), nil
}

func NewUserService() Services {
	return &UserServiceProvider{}
}

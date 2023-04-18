package repository

import (
	"gorm.io/gorm"
)

type Meetup struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"  binding:"required"`
	UserID      uint   `json:"user_id"`
}

//func NewBlogEntity(title string, body string) (*BlogEntity, error) {
//	if utf8.RuneCountInString(title) <= 2 {
//		err := service.MyError{Message: "タイトルを入力してください"}
//		return &BlogEntity{}, err
//	}
//	if body == "" {
//		err := service.MyError{Message: "本文を入力してください"}
//		return &BlogEntity{}, err
//	}
//	return &BlogEntity{Title: title, Body: body}, nil
//}
//
//func (blog *BlogEntity) Create() error {
//	result := db.Create(blog)
//	if result.Error != nil {
//		return result.Error
//	}
//	return nil
//}
//
//func (blog *BlogEntity) FindByTitle() ([]*BlogEntity, error) {
//	var blogs []*BlogEntity
//	result := db.Where("title LIKE ?", "%"+blog.Title+"%").Find(&blogs)
//	if result.Error != nil {
//		return []*BlogEntity{}, result.Error
//	}
//	return blogs, nil
//}
//
//func (blog *BlogEntity) FindById() ([]*BlogEntity, error) {
//	var blogs []*BlogEntity
//	result := db.Where("id = ?", blog.ID).Find(&blogs)
//	if result.Error != nil {
//		return []*BlogEntity{}, result.Error
//	}
//	return blogs, nil
//}

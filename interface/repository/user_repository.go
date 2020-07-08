package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/faozimipa/go-clean/domain/model"

)

type userRepository struct {
	db *gorm.DB
}

//UserRepository interface 
type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}

//NewUserRepository set new
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

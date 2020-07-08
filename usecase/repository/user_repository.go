package repository

import "github.com/faozimipa/go-clean/domain/model"

//UserRepository interface
type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
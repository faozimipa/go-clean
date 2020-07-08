package presenter

import "github.com/faozimipa/go-clean/domain/model"

//UserPresenter interface
type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}

package presenter

import "github.com/faozimipa/go-clean/domain/model"

type userPresenter struct {
}

//UserPresenter interface 
type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
}

//NewUserPresenter set new
func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}

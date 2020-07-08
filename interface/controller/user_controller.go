package controller

import (
	"net/http"

	"github.com/faozimipa/go-clean/domain/model"
	"github.com/faozimipa/go-clean/usecase/interactor"

)

type userController struct {
	userInteractor interactor.UserInteractor
}

//UserController interface
type UserController interface {
	GetUsers(c Context) error
}

//NewUserController func 
func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

//GetUsers get all users
func (uc *userController) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

package registry

import (
	"github.com/faozimipa/go-clean/interface/controller"
	ip "github.com/faozimipa/go-clean/interface/presenter"
	ir "github.com/faozimipa/go-clean/interface/repository"
	"github.com/faozimipa/go-clean/usecase/interactor"
	up "github.com/faozimipa/go-clean/usecase/presenter"
	ur "github.com/faozimipa/go-clean/usecase/repository"

)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}

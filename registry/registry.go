package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/faozimipa/go-clean/interface/controller"

)

type registry struct {
	db *gorm.DB
}

//Registry interface
type Registry interface {
	NewAppController() controller.AppController
}

//NewRegistry func
func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewUserController()
}

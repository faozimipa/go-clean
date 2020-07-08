package main


import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/faozimipa/go-clean/config"
	"github.com/faozimipa/go-clean/infrastructure/datastore"
	"github.com/faozimipa/go-clean/infrastructure/router"
	"github.com/faozimipa/go-clean/registry"

)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
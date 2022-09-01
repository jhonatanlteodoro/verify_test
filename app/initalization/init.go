package initalization

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify_test/app/db"
	"github.com/jhonatanlteodoro/verify_test/app/models"
	"github.com/jhonatanlteodoro/verify_test/app/routes"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (a *App) InitilizeDB() {
	waitSecondsCaseError := 5
	retry_case_error := 5
	conn, err := db.GetConnection(waitSecondsCaseError, retry_case_error)

	if err != nil {
		panic(err)
	}
	a.DB = conn
}

func (a *App) MakeMigrations() {
	models.RunMigrations(a.DB)
}

func (a *App) InitilizeRoutes() {
	a.Router = mux.NewRouter()
}

func (a *App) RegistryRoutes() {
	routes.RegistryRoutes(a.Router, a.DB)
}

func (a *App) Initilize() {
	a.InitilizeDB()
	a.MakeMigrations()
	a.InitilizeRoutes()
	a.RegistryRoutes()
}

func (a *App) Run(host string, port string) {
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), a.Router),
	)
}

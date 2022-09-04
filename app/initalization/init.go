package initalization

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify_test/app/models"
	"github.com/jhonatanlteodoro/verify_test/app/mysql_connector"
	"github.com/jhonatanlteodoro/verify_test/app/routes"
	"github.com/jhonatanlteodoro/verify_test/app/sqlite_connector"
	"gorm.io/gorm"
)

type App struct {
	DB        *gorm.DB
	Router    *mux.Router
	DB_DRIVER string
}

func (a *App) InitilizeDB() {
	waitSecondsCaseError := 5
	retry_case_error := 5

	var conn *gorm.DB
	var err error

	if a.DB_DRIVER == "mysql" {
		conn, err = mysql_connector.GetConnection(waitSecondsCaseError, retry_case_error)
		log.Println("using mysql database")
	} else {
		conn, err = sqlite_connector.GetConnection(waitSecondsCaseError, retry_case_error)
		log.Println("using sqlite database")
	}

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
	a.DB_DRIVER = "mysql"
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

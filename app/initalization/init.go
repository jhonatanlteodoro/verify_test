package initalization

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify_test/app/models"
	"github.com/jhonatanlteodoro/verify_test/app/mysql_connector"
	"github.com/jhonatanlteodoro/verify_test/app/routes"
	"github.com/jhonatanlteodoro/verify_test/app/sqlite_connector"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type App struct {
	DB              *gorm.DB
	Router          *mux.Router
	DB_DRIVER       string
	MYSQL_USERNAME  string
	MYSQL_PASSWORD  string
	MYSQL_DB_NAME   string
	MYSQL_HOST      string
	MYSQL_HOST_PORT string
	SQLITE_FILENAME string
}

func (a *App) LoadEnvVars() {
	dotenvFilePath, envfileErr := filepath.Abs("./.env")
	if envfileErr != nil {
		log.Fatal(envfileErr)
	}

	log.Printf("Loading ENV file from: %s\n", dotenvFilePath)

	err := godotenv.Load(dotenvFilePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a.DB_DRIVER = os.Getenv("DB_IN_USE")
	a.MYSQL_DB_NAME = os.Getenv("MYSQL_DB_NAME")
	a.MYSQL_HOST = os.Getenv("MYSQL_HOST")
	a.MYSQL_USERNAME = os.Getenv("MYSQL_USERNAME")
	a.MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	a.MYSQL_HOST_PORT = os.Getenv("MYSQL_HOST_PORT")
	a.SQLITE_FILENAME = os.Getenv("SQLITE_FILENAME")
}

func (a *App) loadURI() string {
	mysqlUri := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		a.MYSQL_USERNAME, a.MYSQL_PASSWORD,
		a.MYSQL_HOST, a.MYSQL_HOST_PORT,
		a.MYSQL_DB_NAME,
	)
	return mysqlUri
}

func (a *App) LoadSqliteFilePath() string {
	sqliteFile, err := filepath.Abs(fmt.Sprintf("./%s", a.SQLITE_FILENAME))
	if err != nil {
		log.Println("Fail while load sqlite path")
		log.Fatal(err)
	}
	return sqliteFile
}

func (a *App) InitilizeDB() {
	waitSecondsCaseError := 5
	retry_case_error := 5

	var conn *gorm.DB
	var err error

	if a.DB_DRIVER == "mysql" {
		uri := a.loadURI()
		conn, err = mysql_connector.GetConnection(uri, waitSecondsCaseError, retry_case_error)
		log.Println("using mysql database")
	} else {
		sqliteFile := a.LoadSqliteFilePath()
		conn, err = sqlite_connector.GetConnection(sqliteFile, waitSecondsCaseError, retry_case_error)
		log.Println("using sqlite database")
	}

	if err != nil {
		panic(err)
	}

	a.DB = conn
}

func (a *App) MakeMigrations() {
	models.RunMigrations(a.DB)
	log.Println("migrations completed!")
}

func (a *App) InitilizeRoutes() {
	a.Router = mux.NewRouter()
}

func (a *App) RegistryRoutes() {
	routes.RegistryRoutes(a.Router, a.DB)
	log.Println("routers registred!")
}

func (a *App) Initilize() {
	a.LoadEnvVars()
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

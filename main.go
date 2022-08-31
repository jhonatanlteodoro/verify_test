package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify_test/routes"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initilize() {

	a.Router = mux.NewRouter()
	routes.RegistryRoutes(a.Router)
}

func (a *App) Run(host string, port string) {
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), a.Router),
	)
}

func main() {

	a := App{}
	a.Initilize()
	a.Run("", "8080")
}

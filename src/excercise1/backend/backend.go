package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
)

type App struct {
	DB   	*sql.DB
	Port 	string
	Router 	*mux.Router
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func (a *App) Initialize() {
	DB, err := sql.Open("sqlite3", "../../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	a.DB = DB
	a.Router = mux.NewRouter()
	a.initializeRouter()
}

func (a *App) initializeRouter() {
	a.Router.HandleFunc("/", helloWorld)
}

func (a *App) Run() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started and listening on port ", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
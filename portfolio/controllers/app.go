package controllers

import (
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeControllers()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (a *App) initializeControllers() {
	a.Router.HandleFunc("/", a.welcome).Methods("GET")
}

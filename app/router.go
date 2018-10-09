package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/semcelik/go-mongo-example/app/handler"
)

func (a *App) Initialize() {
	fmt.Println("Initializing routes")
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/users", a.GetUsers).Methods("GET")
	a.Router.HandleFunc("/users/{id}", a.GetUser).Methods("GET")
}

func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetUsersHandler(a.Database, w, r)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUserHandler(a.Database, w, r)
}

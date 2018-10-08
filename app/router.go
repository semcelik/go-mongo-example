package app

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/semcelik/go-mongo-example/app/handler"
)

func (a *App) Initialize() {
	fmt.Println("Initializing routes")
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/users", handler.GetUsersHandler).Methods("GET")
}

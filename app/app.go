package app

import (
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type App struct {
	Router *mux.Router
	Client *mongo.Client
}

package app

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func (a *App) InitializeDB() {
	fmt.Println("Initializing db")
	client, err := mongo.NewClient("mongodb://127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	a.Database = client.Database("userdb")
}

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
	a.Client = client
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("userdb").Collection("user")
	res, err := collection.InsertOne(context.Background(), map[string]string{"hello": "world"})
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID
	fmt.Printf("id: %d \n", id)
	fmt.Println("res: ", res)
}

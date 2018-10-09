package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/semcelik/go-mongo-example/app/model"

	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	USER = "user"
)

func GetUsersHandler(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetUsersHandler")
	userCollection := db.Collection(USER)

	cur, err := userCollection.Find(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	var users []model.User

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		user := model.User{}
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
	return
}

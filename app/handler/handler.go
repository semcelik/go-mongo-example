package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/semcelik/go-mongo-example/app/model"
)

const (
	USER = "user"
)

//GetUsersHandler returns all users
func GetUsersHandler(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetUsersHandler")
	userCollection := db.Collection(USER)

	cur, err := userCollection.Find(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	var users []model.User

	defer cur.Close(context.Background())

	elem := bson.NewDocument()

	for cur.Next(context.Background()) {
		elem.Reset()
		user := model.User{}
		err := cur.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}
		user.BsonToUser(elem)
		fmt.Println(user)
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
	return
}

//GetUserHandler returns single user by id
func GetUserHandler(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetUserHandler")
	params := mux.Vars(r)

	userCollection := db.Collection(USER)
	objectID, err := objectid.FromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	elem := bson.NewDocument()
	idDoc := bson.NewDocument(bson.EC.ObjectID("_id", objectID))

	userCollection.FindOne(context.Background(), idDoc).Decode(elem)

	user := model.User{}
	user.BsonToUser(elem)
	json.NewEncoder(w).Encode(user)
	return
}

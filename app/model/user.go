package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
)

type User struct {
	ID       string  `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	LastName string  `json:"lastName,omitempty"`
	Age      float64 `json:"age,omitempty"`
}

//BsonToUser is a workaround for return ObjectID as string
func (u *User) BsonToUser(bson *bson.Document) {
	u.ID = bson.Lookup("_id").ObjectID().Hex()
	u.Name = bson.Lookup("name").StringValue()
	u.LastName = bson.Lookup("lastName").StringValue()
	u.Age = bson.Lookup("age").Double()
}

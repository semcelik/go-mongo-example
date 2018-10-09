package model

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type User struct {
	ID      objectid.ObjectID `bson:"_id" json:"id,omitempty"`
	Name    string            `bson:"name" json:"name,omitempty"`
	Lastname string            `bson:"lastname" json:"hello,omitempty"`
	Age     int               `bson:"age" json:"age,omitempty"`
}

package user

import "gopkg.in/mgo.v2/bson"

type User struct {
	Uid      bson.ObjectId
	Username string
	Password string
	Email    string
}
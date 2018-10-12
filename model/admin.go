package model

import (
	"../db"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Uid      bson.ObjectId
	Username string
	Password string
	Email    string
}

func AddUser(username string, password string, email string) error {
	return db.AdminClient.Users.Insert(&User{bson.NewObjectId(), username,  password,  email})
}

func VerifyUser(username string, password string) (error, User) {
	result := User{}
	err := db.AdminClient.Users.Find(bson.M{"username" : username, "password" : password}).One(&result)

	return err, result
}
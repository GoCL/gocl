package user

import (
	"gocl/db"
	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	Database *db.GoclDatabase
}

func (u *UserModel) Register(username string, password string, email string) error {
	return u.Database.Users.Insert(&User{bson.NewObjectId(), username,  password,  email})
}

func (u *UserModel) Login(username string, password string) (User, error) {
	result := User{}
	err := u.Database.Users.Find(bson.M{"username" : username, "password" : password}).One(&result)

	return result, err
}

var modelPtr *UserModel = nil

func New(databasePtr *db.GoclDatabase) *UserModel {
	if (databasePtr == nil) {
		modelPtr = &UserModel{databasePtr}
	}

	return modelPtr
}

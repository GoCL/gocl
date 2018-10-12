package model

import (
	"gocl/db"
	"gocl/model/user"
)

var (
	UserMode
)

func GetUserModel(databasePtr *db.GoclDatabase) *user.UserModel {
	if (databasePtr == nil) {
		modelPtr = &user.UserModel{databasePtr}
	}

	return modelPtr
}
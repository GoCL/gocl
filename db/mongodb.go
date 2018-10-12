package db

import (
	"gopkg.in/mgo.v2"
	"gocl/config"
)

type GoclDatabase struct {
	Users * mgo.Collection
	Tokens * mgo.Collection
}

func CreateDataBase()(GoclDatabase, error) {
	var database GoclDatabase

	mongoUrl := "mongodb://" + config.DBConf.Url + ":" + config.DBConf.Port
	session, err := mgo.Dial(mongoUrl)

	session.SetMode(mgo.Monotonic, true)
	oauthDB := session.DB("go_admin")
	database.Users = oauthDB.C("users")
	database.Tokens = oauthDB.C("tokens")

	return database, err
}

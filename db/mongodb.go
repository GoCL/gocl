package db

import (
	"gopkg.in/mgo.v2"
	"../conf"
)

type AdminType struct {
	Users * mgo.Collection
	Tokens * mgo.Collection
}

var AdminClient AdminType

func InitMongodb() bool {
	mongoUrl := "mongodb://" + conf.DBConf.Url + ":" + conf.DBConf.Port
	session, err := mgo.Dial(mongoUrl)

	if (err == nil) {
		session.SetMode(mgo.Monotonic, true)
		oauthDB := session.DB("go_admin")
		AdminClient.Users = oauthDB.C("users")
		AdminClient.Tokens = oauthDB.C("tokens")
		return true
	}

	return  false
}

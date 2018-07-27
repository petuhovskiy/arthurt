package main

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var (
	mongoAddr = "localhost"
	dbName    = "rwlist"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
}

func findUserByUsername(username string) (User, error) {
	var user User
	session, err := mgo.Dial("localhost")
	if err != nil {
		return user, err
	}
	defer session.Close()
	query := session.DB(dbName).C("users").Find(bson.M{"username": username})
	err = query.One(&user)
	return user, err
}

func insertNewUser(user *User) error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return err
	}
	defer session.Close()
	return session.DB(dbName).C("users").Insert(user)
}

func init() {
	mongo, err := mgo.Dial(mongoAddr)
	if err != nil {
		log.Fatal(err)
	}

	buildInfo, err := mongo.BuildInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", buildInfo)
}

package db

import (
	mgo "gopkg.in/mgo.v2"
)

func connect() (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("markdown").C("glossary")
	return c, session
}

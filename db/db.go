package db

import (
	mgo "gopkg.in/mgo.v2"
)

func connect(collectionName string) (*mgo.Session, *mgo.Collection) {
	// サーバにデータを保存するため、localhostという名前でセッションを作成
	session, _ := mgo.Dial("localhost")

	db := session.DB("mini-api-server")
	collection := db.C(collectionName)
	return session, collection
}

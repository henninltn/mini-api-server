package db

import (
	mgo "gopkg.in/mgo.v2"
)

func connect(collectionName string) (*mgo.Session, *mgo.Collection) {
	// サーバにデータを保存するため、localhostという名前でセッションを作成
	session, _ := mgo.Dial("localhost")

	// データベースmini-api-serverに接続
	db := session.DB("mini-api-server")
	// mini-api-serverの受け取った名前のコレクションに接続
	collection := db.C(collectionName)
	return session, collection
}

package db

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

// データベースのコレクションthreadのドキュメントに対応した構造体
type Thread struct {
	// bson はデータベースのコレクションのドキュメントに変換する際、される際に使用する
	// json はJSONデータに変換する際、される際に使用する
	ID    bson.ObjectId `bson:"_id"   json:"id"`
	Title string        `bson:"title" json:"title"`
}

type Response struct {
	ID       bson.ObjectId `bson:"id"   json:"id"`
	ThreadID bson.ObjectId `bson:"id"   json:"id"`
	Name     string        `bson:"name" json:"name"`
	Body     string        `bson:"body" json:"body"`
}

// 構造体のデータが仕様に沿っているか確認する
func (thread *Thread) isValid() error {
	// 仕様: タイトルはから文字ではない
	if thread.Title == "" {
		return errors.New("InvalidMemberError at Thread.Title")
	}
	return nil
}

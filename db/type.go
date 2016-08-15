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

// データベースのコレクションresponseに対応した構造体
type Response struct {
	ID bson.ObjectId `bson:"_id"   json:"id"`
	// 対応付けられたthreadのid
	ThreadID bson.ObjectId `bson:"thread_id" json:"thread_id"`
	Name     string        `bson:"name"      json:"name"`
	Body     string        `bson:"body"      json:"body"`
}

// 構造体のデータが仕様に沿っているか確認する
func (thread *Thread) isValid() error {
	// 仕様: タイトルはから文字ではない
	if thread.Title == "" {
		return errors.New("InvalidMemberError at Thread.Title")
	}
	return nil
}

// 構造体のデータが仕様に沿っているか確認する
func (response *Response) isValid() error {
	// 仕様: 対応するスレッドへのIDはから文字ではない
	if response.ThreadID == "" {
		return errors.New("InvalidMemberError at Response.ThreadID")
	}
	// 仕様: 名前が空文字なら"名無し"にする
	if response.Name == "" {
		response.Name = "名無し"
	}
	return nil
}

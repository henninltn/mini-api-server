package db

import "gopkg.in/mgo.v2/bson"

// データベースのコレクションthreadに接続するための関数をまとめるための構造体
type ThreadMapper struct{}

// データベースのコレクションthreadに引数のThread型の構造体を変換して追加する
func (_ *ThreadMapper) Insert(thread *Thread) error {
	// 構造体のデータが仕様に沿っているか確認する
	// db/type.go に定義されている
	if error := thread.isValid(); error != nil {
		return error
	}

	// セッションの開始とデータベースのコレクションthreadへの接続
	session, collection := connect("thread")
	// この関数が終わる時に必ず実行する
	// セッションを閉じる
	defer session.Close()

	// 構造体をドキュメントに変換して接続したコレクションに追加
	error := collection.Insert(thread)
	return error
}

// データベースのコレクションthreadから引数のThread型の構造体のidに一致するドキュメントを取得して、そのデータを構造体の内容で更新する
func (_ *ThreadMapper) Update(thread *Thread) error {
	// 構造体が仕様に沿っているか確認する
	if error := thread.isValid(); error != nil {
		return error
	}

	// セッションの開始とデータベースのコレクションthreadへの接続
	session, collection := connect("thread")
	// この関数が終わる時に必ず実行する
	// セッションを閉じる
	defer session.Close()

	// 接続したコレクションの構造体のidに一致するドキュメントを構造体をドキュメントに変換して、そのデータを構造体を変換したデータの内容で更新
	error := collection.UpdateId(thread.ID, thread)
	return error
}

// データベースのコレクションthreadから引数のThread型の構造体のidに一致するドキュメントを削除する
func (_ *ThreadMapper) Delete(thread *Thread) error {
	// セッションの開始とデータベースのコレクションthreadへの接続
	session, collection := connect("thread")
	// この関数が終わる時に必ず実行する
	// セッションを閉じる
	defer session.Close()

	// 接続したコレクションの構造体のidに一致するドキュメントを削除
	error := collection.RemoveId(thread.ID)
	return error
}

// データベースのコレクションthreadから引数のThread型の構造体のidに一致するドキュメントを取得して、Thread型の構造体に変換して返す
func (_ *ThreadMapper) Find(objectId bson.ObjectId) (*Thread, error) {
	// セッションの開始とデータベースのコレクションthreadへの接続
	session, collection := connect("thread")
	// この関数が終わる時に必ず実行する
	// セッションを閉じる
	defer session.Close()

	// 接続したコレクションの構造体のidに一致するドキュメントを取得して、Thread型の構造体に変換して変数に代入する
	responseThread := new(Thread)
	error := collection.FindId(objectId).One(responseThread)
	return responseThread, error
}

// データベースのコレクションthreadの全てのドキュメントを取得し、Thread型の構造体に変換して、全体をスライスとして返す
func (_ *ThreadMapper) FindAll() (*[]Thread, error) {
	// セッションの開始とデータベースのコレクションthreadへの接続
	session, collection := connect("thread")
	// この関数が終わる時に必ず実行する
	// セッションを閉じる
	defer session.Close()

	// 接続したコレクションのドキュメントを全て取得して、Thread型の構造体に変換して、全体をスライス(のポインタ)として変数に代入する
	responseThreads := new([]Thread)
	error := collection.Find(nil).All(responseThreads)
	return responseThreads, error
}

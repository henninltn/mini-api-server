package handler

import (
	"../db"
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// スレッドに関するリクエストを受け付けるパスを定義する
func ThreadRoutes(router *gin.Engine) {
	/* 下で定義されている、レスポンスとしてJSONを返す関数をそれぞれのリクエストパラメータに登録する */

	router.GET("/thread", getAllThreads)

	// /thread/ 以降をパラメータとして受け取れる
	router.GET("/thread/:id", getThread)

	router.POST("/thread", createThread)

	router.PATCH("/thread", updateThread)

	router.DELETE("/thread", deleteThread)
}

// データベースのコレクションthreadから全てのドキュメントを取得してJSONに変換したものをレスポンスとして返す
func getAllThreads(context *gin.Context) {
	var threadMapper db.ThreadMapper
	// 接続したデータベースのコレクションから全てのドキュメントを取得してThread型の構造体に変換(全て取得しているのでスライスになっている)
	// db/thread.go に定義されている
	responseThreads, error := threadMapper.FindAll()
	if error != nil {
		// ThreadMapper.FindAllに失敗 -> コレクションthreadにはドキュメントが１つも存在しない
		// handler/type.go で定義されている関数
		// status code: 404 Not Found
		context.JSON(404, GetErrorMessage(error))
		return
	}

	// 取得した構造体をJSONに変換してレスポンスする
	// status code: 200 OK
	context.JSON(200, responseThreads)
}

/* データベースのコレクションthreadからリクエストパスで指定されたidに一致するドキュメントを取得してJSONに変換したものをレスポンスとして返す */
func getThread(context *gin.Context) {
	// リクエストパスからパラメータ取得
	stringId := context.Param("id")

	// 文字列のidがデータベースのObjectIdに変換できるか確認して変換
	if !bson.IsObjectIdHex(stringId) {
		// status coede: 400 Bad Request
		context.JSON(400, errors.New("Invalid id"))
		return
	}
	id := bson.ObjectIdHex(stringId)

	var threadMapper db.ThreadMapper
	// 接続したデータベースのコレクションからidが一致するドキュメントを取得してThread型の構造体に変換
	// db/thread.go に定義されている
	responseThread, error := threadMapper.Find(id)
	if error != nil {
		// ThreadMapper.Findに失敗 -> コレクションthreadには指定されたidのドキュメントがない
		// status code: 404 Not Found
		context.JSON(404, GetErrorMessage(error))
		return
	}

	// 取得した構造体をJSONに変換してレスポンスする
	// status code: 200 OK
	context.JSON(200, responseThread)
}

// リクエストのJSONデータを変換してデータベースに追加する
func createThread(context *gin.Context) {
	// リクエストのJSONデータを取得してThread型の構造体に変換
	requestThread := new(db.Thread)
	if error := context.BindJSON(requestThread); error != nil {
		// JSONの変数へのバインド(JSONをThread型の構造体に変換して、変数に代入する)に失敗 -> JSONのデータ形式が間違っている
		// status code: 400 Bad Request
		context.JSON(400, GetErrorMessage(error))
		return
	}
	// IDを生成(リクエストデータにIDは含まれない)
	requestThread.ID = bson.NewObjectId()

	var threadMapper db.ThreadMapper
	// 構造体をドキュメントに変換してデータベースに追加する
	error := threadMapper.Insert(requestThread)
	if error != nil {
		// ThreadMapper.Insertに失敗 -> isValidで失敗したか、DBのInsertに失敗したか
		// どちらにしろJSONのデータ形式が間違っている
		// status code: 400 Bad Request
		context.JSON(400, GetErrorMessage(error))
		return
	}

	// JSONデータはレスポンスせず、作成されたことのみをステータスコードで伝える
	// status code: 201 Created
	context.Status(201)
}

// データベースのコレクションthreadからリクエストのJSONデータのidに一致するドキュメントを取得して、そのデータをリクエストデータのJSONの内容で更新する
func updateThread(context *gin.Context) {
	// リクエストのJSONデータを取得してThread型の構造体に変換
	requestThread := new(db.Thread)
	if error := context.BindJSON(requestThread); error != nil {
		// JSONの変数へのバインド(JSONをThread型の構造体に変換して、変数に代入する)に失敗 -> JSONのデータ形式が間違っている
		// status code: 400 Bad Request
		context.JSON(400, GetErrorMessage(error))
		return
	}

	var threadMapper db.ThreadMapper
	// 構造体のidに一致するデータベースのコレクションthreadのドキュメントを、構造体のデータで内容を更新する
	error := threadMapper.Update(requestThread)
	if error != nil {
		// ThreadMapper.Updateに失敗 -> isValidで失敗したか、DBのInsertに失敗したか
		// どちらにしろJSONのデータ形式が間違っている
		// status code: 400 Bad Request
		context.JSON(400, GetErrorMessage(error))
		return
	}

	// 成功すればリクエストされたJSONをそのまま返す
	// status code: 200 OK
	context.JSON(200, requestThread)
}

// データベースのコレクションthreadからリクエストのJSONデータのidに一致するドキュメントを削除する
func deleteThread(context *gin.Context) {
	// リクエストのJSONデータを取得してThread型の構造体に変換
	requestThread := new(db.Thread)
	if error := context.BindJSON(requestThread); error != nil {
		// JSONの変数へのバインド(JSONをThread型の構造体に変換して、変数に代入する)に失敗 -> JSONのデータ形式が間違っている
		// status code: 400 Bad Request
		context.JSON(400, GetErrorMessage(error))
		return
	}

	var threadMapper db.ThreadMapper
	// 構造体のidに一致するデータベースのコレクションthreadのドキュメントを削除する
	error := threadMapper.Delete(requestThread)
	if error != nil {
		// ThreadMapper.Deleteに失敗 -> DBのInsertに失敗 -> 構造体のidに一致するドキュメントがコレクションthreadにない
		// status code: 404 not found
		context.JSON(404, GetErrorMessage(error))
		return
	}

	// 削除に成功した場合、JSONデータはレスポンスせず削除の成功のみをステータスコードで伝える
	// status code: 204 No Content
	context.Status(204)
}

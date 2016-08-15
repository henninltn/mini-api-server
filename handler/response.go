package handler

import (
	"../db"
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func ResponseRoutes(router *gin.Engine) {
	router.GET("/response/:id", getAllResponse)

	router.POST("/response", makeResponse)

	router.PATCH("/response", updateResponse)

	router.DELETE("/response", deleteResponse)
}

// リクエストパラメータのidに一致するthread_idを持つドキュメントを全て取得して、JSONに変換してレスポンスする
func getAllResponse(context *gin.Context) {
	var stringId = context.Param("id")

	if !bson.IsObjectIdHex(stringId) {
		// status coede: 400 Bad Request
		context.JSON(400, errors.New("Invalid id"))
		return
	}
	objectId := bson.ObjectIdHex(stringId)

	var responseMapper db.ResponseMapper
	responseThreads, error := responseMapper.FindAll(objectId)
	if error != nil {
		// status code: 404 Not Found
		context.JSON(404, GetErrorMessage(error))
		return
	}

	// status code: 200 OK
	context.JSON(200, responseThreads)
}

// リクエストのJSONデータを変換してコレクションに追加する
func makeResponse(context *gin.Context) {
	requestResponse := new(db.Response)
	if error := context.BindJSON(requestResponse); error != nil {
		// status code: 400 Bad Request
		context.JSON(400, GetErrorMessage(error))
		return
	}
	requestResponse.ID = bson.NewObjectId()

	var responseMapper db.ResponseMapper
	error := responseMapper.Insert(requestResponse)
	if error != nil {
		// status code: 400 Bad Request
		context.JSON(400, GetErrorMessage(error))
		return
	}

	context.Status(201)
}

// リクエストのJSONデータのidに一致するドキュメントを、そのJSONデータの内容で更新する
func updateResponse(context *gin.Context) {
	requestResponse := new(db.Response)
	if error := context.BindJSON(requestResponse); error != nil {
		// status code: 400 Bad Request
		context.JSON(400, GetErrorMessage(error))
		return
	}

	var responseMapper db.ResponseMapper
	error := responseMapper.Update(requestResponse)
	if error != nil {
		// status code: 400 Bad Request
		context.JSON(400, GetErrorMessage(error))
		return
	}

	// status code: 200 OK
	context.JSON(200, requestResponse)
}

// リクエストのJSONデータのidに一致するドキュメントを削除する
func deleteResponse(context *gin.Context) {
	requestResponse := new(db.Response)
	if error := context.BindJSON(requestResponse); error != nil {
		// status code: 400 Bad Request
		context.JSON(400, GetErrorMessage(error))
		return
	}

	var responseMapper db.ResponseMapper
	error := responseMapper.Delete(requestResponse)
	if error != nil {
		// status code: 404 Not Found
		context.JSON(404, GetErrorMessage(error))
		return
	}

	// status code: No Content
	context.Status(204)
}

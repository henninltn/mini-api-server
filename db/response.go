package db

import "gopkg.in/mgo.v2/bson"

type ResponseMapper struct{}

// 構造体をドキュメントに変換してコレクションに追加する
func (_ *ResponseMapper) Insert(response *Response) error {
	if error := response.isValid(); error != nil {
		return error
	}

	session, collection := connect("response")
	defer session.Close()

	error := collection.Insert(response)
	return error
}

// 構造体のidに一致するドキュメントをコレクションresponseから取得して、構造体の内容でデータを更新する
func (_ *ResponseMapper) Update(response *Response) error {
	if error := response.isValid(); error != nil {
		return error
	}

	session, collection := connect("response")
	defer session.Close()

	error := collection.UpdateId(response.ID, response)
	return error
}

// 構造体のidに一致するドキュメントをコレクションresponseから取得して、削除する
func (_ *ResponseMapper) Delete(response *Response) error {
	session, collection := connect("response")
	defer session.Close()

	error := collection.RemoveId(response.ID)
	return error
}

// idに一致するthread_idを持ったドキュメントを全て取得して、全体をスライス(のポインタ)として返す
func (_ *ResponseMapper) FindAll(objectId bson.ObjectId) (*[]Response, error) {
	session, collection := connect("response")
	defer session.Close()

	responseThreads := new([]Response)
	error := collection.Find(bson.M{"thread_id": objectId}).All(responseThreads)
	return responseThreads, error
}

package handler

import "fmt"

// エラー時にレスポンスするJSON用の構造体
type errorMessage struct {
	// json はJSONデータに変換する際、される際に使用する
	Error string `json:"error"`
}

// error型の変数をJSONに変換するための構造体errorMessageに変換する
func GetErrorMessage(error error) errorMessage {
	// fmt.Sprint は画面に出力するのではなく、変数の中に文字列として出力する
	return errorMessage{fmt.Sprint(error)}
}

package service

import (
	"encoding/json"
	"gitstats/src/model"
	"net/http"
)

const SUCCESS = 0
const ERROR = -1

func Success(data interface{}) model.Response {
	return model.Response{
		Code: SUCCESS,
		Msg:  "success",
		Data: data,
	}
}

func Failure(code int, msg string) model.Response {
	return model.Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func Error(msg string) model.Response {
	return Failure(ERROR, msg)
}

func ToJson(result model.Response) []byte {
	jsonData, err := json.Marshal(result)
	if err != nil {
		return nil
	}
	return jsonData
}

func WriteJson(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func SetAccess(w http.ResponseWriter) {
	// 设置允许跨域的域名
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 设置允许的请求方法
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 设置预检请求的缓存时间
	w.Header().Set("Access-Control-Max-Age", "3600")
	// 设置允许的请求头
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

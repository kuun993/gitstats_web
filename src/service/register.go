package service

import (
	"net/http"
)

// RegisterHttpHandle /**	注册http请求处理函数
func RegisterHttpHandle(url string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(url, handler)
}

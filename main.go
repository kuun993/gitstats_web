package main

import (
	"fmt"
	"gitstats/src/service"
	_ "io/ioutil"
	"net/http"
	_ "os"
)

func main() {

	// 注册 http 处理器
	service.RegisterHttpHandle("/gitstats", service.HandleGitStats)

	err := http.ListenAndServe(":19877", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

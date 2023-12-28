package service

import (
	"net/http"
)

func HandleGitStats(w http.ResponseWriter, r *http.Request) {
	SetAccess(w)
	// 获取查询参数
	queryParams := r.URL.Query()
	git := queryParams.Get("git")
	if git == "" {
		return
	}

	// 克隆git项目到本地
	name := clone(git)
	if name == "" {
		WriteJson(w, ToJson(Error("clone failed")))
		return
	}

	// 获取git项目的统计信息
	url := gitstats(name)
	if url == "" {
		WriteJson(w, ToJson(Error("gitstats failed")))
		return
	}

	jsonData := ToJson(Success(url))

	WriteJson(w, jsonData)
}

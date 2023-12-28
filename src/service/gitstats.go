package service

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

const NGINX_URL = "http://10.1.1.15:9876/"

const GIT_STATS = "/opt/gitstats/gitstats"

const SRC = "/opt/src/"

const OUTPUT = "/opt/output/"

func clone(git string) string {
	// 切换到src目录
	var err = os.Chdir(SRC)
	if err != nil {
		fmt.Println("not found dir:", err)
		return ""
	}

	// 获取git项目名称
	name := strings.Replace(path.Base(git), path.Ext(git), "", -1)

	// 判断项目是否已经存在
	_, err = os.Stat(SRC + name)
	if err == nil {
		os.Chdir(SRC + name)
		// 更新git项目
		cmd := exec.Command("git", "pull")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error git pull:", err)
			return ""
		}
		return name
	}

	// 克隆git项目到本地
	cmd := exec.Command("git", "clone", git)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error cloning git repo:", err)
		return ""
	}
	return name
}

func gitstats(name string) string {
	// 删除output目录下的项目
	os.RemoveAll(OUTPUT + name)

	// 获取git项目的统计信息
	cmd := exec.Command(GIT_STATS, SRC+name, OUTPUT+name)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error cloning git repo:", err)
		return ""
	}
	return NGINX_URL + path.Base(name) + "/"
}

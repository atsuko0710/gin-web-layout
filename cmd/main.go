package main

import (
	"bluebell/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "project config file path.")
)

func main() {

	// 初始化配置文件
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 初始化数据库

	r := gin.Default()

	r.Run()
}

package main

import (
	"my-singo/conf"
	"my-singo/server"
	"os"
)

// @title 实验项目
// @version 1.0
// @description 陈佳实验项目接口文档
// @host localhost:8080
// @BasePath /api/v1
func main() {

	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()

	// 服务启动端口
	if os.Getenv("SERVER_PORT") != "" {
		r.Run(":" + os.Getenv("SERVER_PORT"))
	} else {
		r.Run(":8080")
	}

}

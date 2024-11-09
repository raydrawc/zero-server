package main

import "github.com/gin-gonic/gin"

func init() {
	// 加载配置

	// 初始化基础库
}

func main() {
	print("hello world")
	HttpServer()
}

func HttpServer() {
	// 启动http服务
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	r.Run(":8080")
}

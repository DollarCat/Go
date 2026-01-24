package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//配置路由
	r.GET("/")
}

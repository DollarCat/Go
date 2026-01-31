package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	//判断用户是否登录

	fmt.Println(time.Now())

	fmt.Println(c.Request.URL)

	c.Set("username", "Alex")

	/*
		定义一个goroutine统计日志  当在中间件或 handler 中启动新的 goroutine 时
		不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）
		- 原始的 c *gin.Context 在请求处理完成后会被回收
		- 如果直接使用 c ，goroutine 可能会访问到已释放的内存
		- 使用 c.Copy() 创建一个 只读副本，确保 goroutine 安全访问
	*/
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}

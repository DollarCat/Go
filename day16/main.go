package main

import (
	"day16/models"
	"day16/routers"
	"fmt"
	"html/template"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-我是一个中间件")
	//调用该请求的剩余处理程序
	c.Next()
	//Abort 是终止的意思， c.Abort() 表示终止调用该请求的剩余处理程序

	fmt.Println("2-我是一个中间件")
	end := time.Now().UnixNano()

	fmt.Println(end - start)
}

func initMiddleware2(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-我是一个中间件---initMiddleware2")
	//调用该请求的剩余处理程序
	c.Next()

	fmt.Println("2-我是一个中间件---initMiddleware2")
	end := time.Now().UnixNano()

	fmt.Println(end - start)
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	//加载模板 放在配置路由前面
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	//配置session中间件

	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret111"))
	//配置session的中间件 store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))
	/*
		//配置session中间件 把session存储到Redis
		store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret111"))
		r.Use(sessions.Sessions("mysession", store))
	*/

	r.GET("/initMiddleware", initMiddleware, initMiddleware2, func(c *gin.Context) {
		fmt.Println("这是一个首页")
		// time.Sleep(time.Second)
		c.String(200, "gin首页")
	})

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	r.Run(":80")
}

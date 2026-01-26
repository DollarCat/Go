package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	//- 第一个参数 sec ：秒数（Unix 时间戳）- 第二个参数 nsec ：纳秒数（用于更精确的时间）
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + "----" + str2
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前
	//模板加载时会解析模板内容，如果函数还没注册，模板引擎会报错找不到函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})
	//加载模板
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")
	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "aaa",
			"msg":   " 我是msg",
			"score": 89,
			"hobby": []string{"吃饭", "睡觉", "写代码"},
			"newsList": []interface{}{
				&Article{
					Title:   "新闻标题111",
					Content: "新闻详情111",
				},
				&Article{
					Title:   "新闻标题222",
					Content: "新闻详情222",
				},
			},
			"testSlice": []string{},
			"news": &Article{
				Title:   "新闻标题",
				Content: "新闻内容",
			},
			"date": 1629423555,
		})
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "这是一个post--主要用于增加数据")
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "这是一个put请求 主要用于编辑数据")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "这是一个DELETE请求 用于删除数据")
	})

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好gin",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好gin--22",
		})
	})

	r.GET("/json3", func(c *gin.Context) {

		/*
			使用指针：
			１、避免值拷贝（性能考虑）：如果结构体很大（有很多字段），值拷贝会消耗更多内存和时间。
			指针只占用 8 字节（64位系统），效率更高。
			２、go惯例，对于较大结构体，通常使用指针传递
		*/
		a := &Article{
			Title:   "我是一个标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSON(200, a)
	})

	//响应Jsonp请求
	// http://localhost:8080/jsonp?callback=xxxx
	// xxxx({"title":"我是一个标题-jsonp","desc":"描述","content":"测试内容"});
	r.GET("/jsonp", func(c *gin.Context) {

		a := &Article{
			Title:   "我是一个标题-jsonp",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSONP(200, a)
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "你好gin 我是一个xml",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻详情",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻页面",
			"news":  news,
		})
	})

	//后台
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})
	r.GET("/admin/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "新闻页面",
		})
	})

	// r.Run()  启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务

	r.Run(":8000") //启动一个web服务
}

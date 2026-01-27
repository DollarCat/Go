package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
同一个结构体可以处理多种数据格式：
- GET请求 → 使用 form 标签匹配URL参数
- POST表单 → 使用 form 标签匹配表单数据
- POST JSON → 使用 json 标签匹配JSON字段
- POST XML → 使用 xml 标签匹配XML元素
*/
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	//加载模板 放在配置路由前面
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	//Get 请求传值
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	//Get 请求传值   id
	r.GET("/article", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"msg": "新闻详情",
			"id":  id,
		})
	})

	//post演示
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/user.html", gin.H{})
	})
	//获取表单post过来的数据
	r.POST("/doAddUser1", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	//获取 GET POST 传递的数据绑定到结构体

	//http://localhost:8080/getUser?username=zhangsan&password=1111

	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}
		/*
			c.ShouldBind自动根据请求类型（GET/POST）和结构体标签，将请求参数绑定到 user 结构体
			通过 结构体标签 来匹配字段，如：`json:"username" form:"username"`
			绑定成功返回 nil ，失败返回错误

		*/
		if err := c.ShouldBind(&user); err == nil {
			fmt.Printf("%#v", user)
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})

	r.POST("/doAddUser2", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})
	//获取 Post Xml 数据
	r.POST("/xml", func(c *gin.Context) {
		article := &Article{}

		xmlSliceData, _ := c.GetRawData() //获取 c.Request.Body 读取请求数据

		fmt.Println(xmlSliceData)

		if err := xml.Unmarshal(xmlSliceData, &article); err == nil {
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}

	})
	/*
		动态路由传值
		- : 表示这是一个动态参数
		- cid 是参数名
		使用接口方式：list/123          list/456
	*/
	r.GET("/list/:cid", func(c *gin.Context) {

		cid := c.Param("cid")
		c.String(200, "%v", cid)

	})

	//后台
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})
	r.Run()

}

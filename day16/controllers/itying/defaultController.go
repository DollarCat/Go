package itying

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	//设置cookie
	//3600表示的是秒
	c.SetCookie("username", "张三", 3600, "/", "localhost", false, true)

	//设置sessions
	session := sessions.Default(c)
	session.Set("sessionName", "天下第一！！！")
	session.Save() //设置session的时候必须调用

	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "我是一个msg",
		"t":   1629788418,
	})
}
func (con DefaultController) News(c *gin.Context) {
	//获取sessions
	session := sessions.Default(c)
	username := session.Get("sessionName")
	c.String(200, "username=%v", username)
}

func (con DefaultController) Shop(c *gin.Context) {
	//获取cookie
	username, _ := c.Cookie("username")
	c.String(200, "username=%v", username)
}

func (con DefaultController) DeleteCookie(c *gin.Context) {
	//删除cookie
	c.SetCookie("username", "张三", -1, "/", "localhost", false, true)
	c.String(200, "删除成功")
}

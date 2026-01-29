package admin

import "github.com/gin-gonic/gin"

type UserController struct {
	BaseController // 嵌入，继承所有方法
}

func (con UserController) Index(c *gin.Context) {
	// c.String(200, "用户列表--")
	con.success(c) // 调用继承的方法
}
func (con UserController) Add(c *gin.Context) {
	c.String(200, "用户列表-add---")
}
func (con UserController) Edit(c *gin.Context) {
	c.String(200, "用户列表-Edit------")
}

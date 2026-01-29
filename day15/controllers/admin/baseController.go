package admin

import "github.com/gin-gonic/gin"

/*
BaseController 是一个基础控制器 ，用于提供 公共方法和属性 ，让其他控制器继承使用。

好处 ：

- 代码复用 ：多个控制器都需要相同的方法（如 success、error），不用重复写
- 统一管理 ：公共逻辑集中在一个地方，便于维护
- 继承机制 ：其他控制器通过嵌入 BaseController 来获得这些方法
*/
type BaseController struct {
}

func (con BaseController) success(c *gin.Context) {
	c.String(200, "成功")
}

func (con BaseController) error(c *gin.Context) {
	c.String(200, "失败")
}

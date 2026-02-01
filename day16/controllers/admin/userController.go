package admin

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	// c.String(200, "用户列表--")
	con.success(c)
}

func (con UserController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
}

func (con UserController) Edit(c *gin.Context) {
	c.String(200, "用户列表-Edit------")
}

func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")

	/*
		//获取单个文件代码：
		file, err := c.FormFile("face")

		// file.Filename 获取文件名称  aaa.jpg   ./static/upload/aaa.jpg
		dst := path.Join("./static/upload", file.Filename)
		if err == nil {
			c.SaveUploadedFile(file, dst)
		}
		// c.String(200, "执行上传")
		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"username": username,
			"dst":      dst,
		})
	*/

	form, _ := c.MultipartForm()
	files := form.File["face[]"]

	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		// 上传文件至指定目录
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": username,
	})
}

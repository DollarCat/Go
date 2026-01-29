package routers

import (
	"day15/controllers/alex"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", alex.DefaultController{}.Index)
		defaultRouters.GET("/news", alex.DefaultController{}.News)
	}
}

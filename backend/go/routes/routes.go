package routes

import (
	"github/wangcham/kuo_daibiao/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine{
	r := gin.Default()


	userGroup := r.Group("user")
	{
		userGroup.POST("/users",controller.CreateUser)
	}
	return r
}
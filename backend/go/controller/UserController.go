package controller

import(
	"github/wangcham/kuo_daibiao/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"github/wangcham/kuo_daibiao/service"
)

func CreateUser(c *gin.Context){
	var user *entity.User
	c.BindJSON(&user)
	err := service.CreateUser(user)

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	 }else {
		c.JSON(http.StatusOK,gin.H{
		   "code":200,
		   "msg":"success",
		   "data":user,
		})
	 }

}
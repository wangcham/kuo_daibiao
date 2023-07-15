package main

import (
	"github/wangcham/kuo_daibiao/dao"
	"github/wangcham/kuo_daibiao/entity"
	"github/wangcham/kuo_daibiao/routes"
)

func main(){
	err := dao.InitMySql()
	if err!=nil{
		panic(err)
	}
	defer dao.Close()
	dao.SqlSession.AutoMigrate(&entity.User{},&entity.AllCourse{},&entity.Course{})
	r := routes.SetRouter()
	r.Run(":8081")
}
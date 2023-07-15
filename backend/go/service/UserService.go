package service

import (
	"github/wangcham/kuo_daibiao/dao"
	"github/wangcham/kuo_daibiao/entity"
)

func CreateUser(user *entity.User) (err error){
	if err := dao.SqlSession.Error;err!=nil{
		return err
	}
	return
}


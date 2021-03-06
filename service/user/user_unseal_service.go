package user

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// 封禁用户的服务
type UserUnsealService struct{}

// 用户封禁函数
func (service *UserUnsealService) Unseal(c *gin.Context) serializer.Response {
	//检测root
	if !funcs.CheckRoot(c) {
		return serializer.CheckNoRight()
	}
	uid := c.Param("uid")
	//获取封禁用户
	user2 := model.User{}
	if err := model.DB.First(&user2, uid).Error; err != nil {
		return serializer.DBErr("", err)
	}
	//更新数据库数据
	user2.Status = model.Active
	if err := model.DB.Save(&user2).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("解封成功", true)
}

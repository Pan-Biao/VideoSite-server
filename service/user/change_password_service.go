package user

import (
	"regexp"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// ChangePasswordService 用户修改密码服务
type ChangePasswordService struct {
	OldPassword string `form:"old_password" json:"old_password"`
	NewPassword string `form:"new_password" json:"new_password"`
}

func (service *ChangePasswordService) Change(c *gin.Context) serializer.Response {
	user := funcs.GetUser(c)

	if re, _ := regexp.MatchString("^[a-zA-Z0-9]{6,16}$", service.OldPassword); !re {
		return serializer.ParamErr("密码格式错误,应为6-16位数字或大小写字母")
	}

	if re, _ := regexp.MatchString("^[a-zA-Z0-9]{6,16}$", service.NewPassword); !re {
		return serializer.ParamErr("密码格式错误,应为6-16位数字或大小写字母")
	}

	//验证密码
	if !user.CheckPassword(service.OldPassword) {
		return serializer.ParamErr("旧密码错误")
	}

	// 加密密码
	if err := user.SetPassword(service.NewPassword); err != nil {
		return serializer.SetErr("密码加密失败", err)
	}

	// 修改密码
	if err := model.DB.Save(&user).Error; err != nil {
		return serializer.SetErr("密码修改失败", err)
	}

	return serializer.ReturnData("密码修改成功", true)
}

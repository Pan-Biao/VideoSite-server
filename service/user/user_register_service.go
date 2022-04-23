package user

import (
	"unicode/utf8"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname"`
	UserName        string `form:"user_name" json:"user_name"`
	Password        string `form:"password" json:"password"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	if utf8.RuneCountInString(service.Nickname) < 2 || utf8.RuneCountInString(service.Nickname) > 8 {
		sr := serializer.ParamErr("昵称长度应为2-8个字")
		return &sr
	}
	if utf8.RuneCountInString(service.UserName) < 6 || utf8.RuneCountInString(service.UserName) > 12 {
		sr := serializer.ParamErr("用户名长度应为6-12位数")
		return &sr
	}
	if utf8.RuneCountInString(service.UserName) < 6 || utf8.RuneCountInString(service.UserName) > 16 {
		sr := serializer.ParamErr("密码长度应为6-16位数")
		return &sr
	}
	if service.PasswordConfirm != service.Password {
		sr := serializer.ParamErr("两次输入的密码不相同")
		return &sr
	}
	count := int64(0)
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		sr := serializer.ParamErr("昵称被占用")
		return &sr
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		sr := serializer.ParamErr("用户名已经注册")
		return &sr
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   model.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.SetErr("密码加密失败", err)
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.SetErr("注册失败", err)
	}

	return serializer.ReturnData("注册成功", serializer.BuildUser(user))
}

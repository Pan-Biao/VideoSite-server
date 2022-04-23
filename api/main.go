package api

import (
	"encoding/json"
	"fmt"
	"log"
	"vodeoWeb/cache"
	"vodeoWeb/conf"
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// CurrentUser 获取当前用户
// func CurrentUser(c *gin.Context) *model.User {
// 	if user, _ := c.Get("user"); user != nil {
// 		if u, ok := user.(*model.User); ok {
// 			return u
// 		}
// 	}
// 	return nil
// }

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	token := c.GetHeader("Authorization")
	if token != "" {
		log.Panicln("使用token获取当前用户")
		if id, _ := cache.GetUserByToken(token); id != "" {
			user := model.User{}
			model.DB.First(&user, id)
			return &user
		}
	} else {
		log.Panicln("未使用token")
		if user, _ := c.Get("user"); user != nil {
			if u, ok := user.(*model.User); ok {
				return u
			}
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.ParamErr(
				fmt.Sprintf("%s%s", field, tag),
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配")
	}

	return serializer.ParamErr("参数错误")
}

package middleware

import (
	"fmt"
	"vodeoWeb/cache"
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uid string
		token := c.GetHeader("Authorization")
		if token != "" {
			fmt.Println("---------------token--------------")
			uid, _ = cache.GetUserByToken(token)
		} else {
			fmt.Println("---------------session--------------")
			session := sessions.Default(c)
			uid, _ = session.Get("user_id").(string)
			fmt.Println("----------", uid)
		}
		if uid != "" {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}

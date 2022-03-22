package server

import (
	"net/http"
	"os"
	"vodeoWeb/api"
	"vodeoWeb/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET"))) //Session保持登录状态
	r.Use(middleware.Cors())                               //跨域
	r.Use(middleware.CurrentUser())                        //用户登录状态

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)
		// 用户注册
		v1.POST("user/register", api.UserRegister)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			//用户接口
			authed.GET("user/me", api.UserMe)            //获取用户数据
			authed.DELETE("user/logout", api.UserLogout) //退出登录
			authed.PUT("user/update", api.UpdateUser)    //修改用户数据
		}

		//分区接口
		v1.GET("sub_areas", api.ListSubArea)
		{
			authed.POST("sub_area", api.CreateSubArea)
			authed.PUT("sub_area/:said", api.UpdateSubArea)
			authed.DELETE("sub_area/:said", api.DeleteSubArea)
		}

		//视频接口
		v1.GET("video/:id", api.ShowVideo)            //视频信息
		v1.POST("videos/search", api.ListSearchVideo) //获取视频列表
		v1.GET("video/play/:vid", api.PlayVideo)      //视频文件
		{
			authed.POST("video", api.CreateVideo)        //创建视频
			authed.PUT("video/:vid", api.UpdateVideo)    //更新视频
			authed.DELETE("video/:vid", api.DeleteVideo) //删除视频
		}

		//关注接口
		v1.GET("follower", api.ListFollower)
		v1.GET("fans", api.ListFans)
		{
			authed.POST("follow/:fid", api.CreateFollow)
			authed.DELETE("follow/:fid", api.DeleteFollow)
		}

		//收藏夹接口
		{
			authed.POST("favorites", api.ListFavorite)
			authed.POST("favorite", api.CreateFavorite)
			authed.PUT("favorite/:fid", api.UpdateFavorite)
			authed.DELETE("favorite/:fid", api.DeleteFavorite)
		}

		//收藏接口
		{
			authed.POST("collections", api.ListCollection)
			authed.POST("collection", api.CreateCollection)
			authed.DELETE("collection/:cid", api.DeleteCollection)
		}

		//用户静态资源
		v1.Static("/assets", "./assets")
		v1.StaticFS("/file", http.Dir("G:/videoResources"))
	}

	return r
}

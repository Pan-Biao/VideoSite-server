package server

import (
	"net/http"
	"vodeoWeb/api"
	"vodeoWeb/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session("BurstMagic"))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	//用户静态资源
	file := r.Group("/")
	file.Static("/assets", "./assets")
	file.StaticFS("file", http.Dir("G:/videoResources"))

	// 路由
	v1 := r.Group("/api/v1")
	//用户登录状态
	{
		v1.POST("ping", api.Ping)
		// 用户注册
		v1.POST("user/register", api.UserRegister)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
		//获取用户信息
		v1.GET("user/:uid", api.UserInformation)
		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			//用户接口
			authed.GET("user/me", api.UserMe)                       //获取用户数据
			authed.POST("user/refresh", api.UserTokenRefresh)       //用户刷新token
			authed.POST("user/change_password", api.ChangePassword) //修改密码
			authed.DELETE("user/logout", api.UserLogout)            //退出登录
			authed.PUT("user/update", api.UpdateUser)               //修改用户数据
		}

		//分区列表接口
		v1.GET("sub_area/list", api.ListSubArea)
		//轮播列表
		v1.GET("carousel/list", api.ListCarousel)

		//视频接口
		v1.GET("video/:vid", api.ShowVideo)        //视频信息
		v1.POST("video/list", api.ListSearchVideo) //获取视频列表
		// v1.POST("video/play/:vid", api.PlayVideo)         //视频文件
		v1.POST("video/play/:vid", api.PlayNumber) //增加视频播放数
		{
			authed.POST("video/create", api.CreateVideo)      //创建视频
			authed.PUT("video/:vid", api.UpdateVideo)         //更新视频
			authed.DELETE("video/:vid", api.DeleteVideo)      //删除视频
			authed.POST("video/like/:vid", api.LikeVideo)     //点赞
			authed.GET("video/islike/:vid", api.IsLikeVideo)  //是否点赞
			authed.POST("video/unlike/:vid", api.UnLikeVideo) //取消点赞
		}
		//评论接口
		v1.POST("comment/list", api.ListComment)
		{
			authed.POST("comment/:vid", api.CreateComment)
			authed.DELETE("comment/:cid", api.DeleteComment)
			authed.POST("comment/like/:cid", api.LikeComment)
			authed.POST("comment/unlike/:cid", api.UnLikeComment)
			authed.GET("comment/islike/list", api.IsLikeComment)
		}

		//关注接口
		v1.GET("follow/follower/:uid", api.ListFollower)
		v1.GET("follow/fans/:uid", api.ListFans)
		{
			authed.POST("follow/:fid", api.CreateFollow)
			authed.DELETE("follow/:fid", api.DeleteFollow)
			authed.GET("follow/isfollow/:uid", api.IsFollow)
		}

		//收藏夹接口
		{
			authed.GET("favorite/list", api.ListFavorite)
			authed.POST("favorite/add", api.CreateFavorite)
			authed.PUT("favorite/:fid", api.UpdateFavorite)
			authed.DELETE("favorite/:fid", api.DeleteFavorite)
		}

		//收藏接口
		{
			authed.GET("collection/list/:fid", api.ListCollection)
			authed.POST("collection/add", api.CreateCollection)
			authed.GET("collection/iscollection/:cid", api.IsCollection)
			authed.DELETE("collection/:cid", api.DeleteCollection)
		}

	}

	// root路由
	root := r.Group("/api/root")
	// 需要登录保护
	root.Use(middleware.AuthRequired())
	//分区接口
	{
		root.POST("sub_area/add", api.CreateSubArea)
		root.PUT("sub_area/:said", api.UpdateSubArea)
		root.DELETE("sub_area/:said", api.DeleteSubArea)
	}
	//轮播图
	{
		root.POST("carousel/add", api.AddCarousel)
		root.DELETE("carousel/:cid", api.DeleteCarousel)
	}
	//用户
	{
		root.POST("user/list", api.UserList)
		root.POST("user/suspend/:uid", api.UserSuspend) //禁用
		root.POST("user/unseal/:uid", api.UserUnseal)   //解封
	}
	//视频
	{
		root.POST("video/list", api.ListRootVideo)        //获取视频列表
		root.POST("video/suspend/:vid", api.VideoSuspend) //禁用
		root.POST("video/unseal/:vid", api.VideoUnseal)   //解封
	}

	return r
}

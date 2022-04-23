package conf

import (
	"vodeoWeb/cache"
	"vodeoWeb/model"
	"vodeoWeb/util"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()
	gin.SetMode("release")

	// 设置日志级别
	util.BuildLogger("info")

	// 读取翻译文件
	// if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
	// 	util.Log().Panic("翻译文件加载失败", err)
	// }

	// 连接数据库
	model.Database("root:qq201100qq@tcp(127.0.0.1:3306)/video_web?charset=utf8&parseTime=True&loc=Local")

	cache.Redis()
}

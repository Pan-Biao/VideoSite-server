package video

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// 视频点赞服务
type IsLikeVideoService struct{}

func (service *IsLikeVideoService) Is(c *gin.Context) serializer.Response {
	//查找对应视频
	vid := c.Param("vid")
	user := funcs.GetUser(c)
	videoLike := model.VideoLike{}

	db := model.DB
	db = db.Where("uid = ?", user.ID)
	db = db.Where("vid = ?", vid)
	if re := funcs.SQLErr(db.First(&videoLike).Error); re != nil {
		return serializer.Response{
			Code: 200,
			Data: false,
		}
	}

	return serializer.Response{
		Code: 200,
		Data: true,
	}
}

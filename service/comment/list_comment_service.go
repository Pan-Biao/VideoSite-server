package comment

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/util"

	"github.com/gin-gonic/gin"
)

// Sorts 排序条件列表
type Sorts struct {
	Sort  string `form:"sort" json:"sort"`
	Field string `form:"field" json:"field"`
}

// 评论列表
type ListCommentService struct {
	Sorts *[]Sorts `form:"sorts" json:"sorts"`
	Vid   int      `form:"vid" json:"vid"`
}

// 评论列表服务
func (service *ListCommentService) List(c *gin.Context) serializer.Response {
	comments := []model.Comment{}
	db := model.DB

	//添加多个排序条件
	if service.Sorts != nil {
		for _, v := range *service.Sorts {
			if v.Field == "" {
				return serializer.ParamErr("排序条件错误")
			} else {
				//拼接字符串
				strs := []string{v.Field, " ", v.Sort}
				db = db.Order(util.Join(strs))
			}
		}
	} else {
		db = db.Order("created_at desc")
	}

	db = db.Where("vid = ?", service.Vid)
	//查询
	if err := db.Find(&comments).Error; err != nil {
		return serializer.DBErr("", err)
	}
	//反回数据
	return serializer.ReturnData("成功", serializer.BuildComments(comments))
}

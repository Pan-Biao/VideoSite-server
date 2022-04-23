package video

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/util"

	"github.com/gin-gonic/gin"
)

// ListRootVideoService 视频列表
type ListRootVideoService struct {
	Number     *int     `form:"number" json:"number"`
	Sorts      []Sorts  `form:"sorts" json:"sorts"`
	Searchs    []string `form:"searchs" json:"searchs"`
	PageNumber int      `form:"page_number" json:"page_number"`
}

// List 视频列表服务
func (service *ListRootVideoService) List(c *gin.Context) serializer.Response {
	videos := []model.Video{}
	db := model.DB
	//总长度
	var total int64
	//每页的长度 不传入默认20 传入0为不限制长度,此时不需要传入当前页PageNumber
	number := 0
	if service.Number == nil {
		number = 20
		db = db.Limit(number)
	} else {
		number = *service.Number
		db = db.Limit(number)
	}
	//添加多个排序条件
	if service.Sorts != nil {
		for _, v := range service.Sorts {
			if v.Field == "" {
				return serializer.ParamErr("排序条件错误")
			} else {
				//拼接字符串
				strs := []string{v.Field, " ", v.Sort}
				db = db.Order(util.Join(strs))
			}
		}
	}

	//关键字查找
	if service.Searchs != nil {
		//循环添加需要查询的关键字
		for _, v := range service.Searchs {
			if v != "" {
				str := util.JoinLike(v)
				//先查找标题
				db = db.Where("title like ?", str)
				//然后查找昵称
				users := []model.User{}
				model.DB.Where("nickname like ?", str).Find(&users)
				for _, u := range users {
					db = db.Or("uid = ?", u.ID)
				}
			}
		}
	}

	//分页 当前页数 无参数时默认为第1页
	//db.Offset() 必须放在find()前面，不然查询语句可能不对
	pageNumber := 1
	if service.PageNumber > 1 && number > 0 {
		pageNumber = service.PageNumber
		db = db.Offset((pageNumber - 1) * number)
	}
	//查询
	if err := db.Find(&videos).Error; err != nil {
		return serializer.DBErr("", err)
	}
	//查询长度
	db = db.Limit(-1).Offset(-1)
	if err := db.Count(&total).Error; err != nil {
		return serializer.DBErr("", err)
	}
	//反回数据
	return serializer.ReturnData("成功", serializer.BuildVideos(videos, pageNumber, number, int(total)))
}

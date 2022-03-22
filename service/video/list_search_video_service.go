package video

import (
	"log"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/util"
)

// Sorts 排序条件列表
type Sorts struct {
	Sort  string `form:"sort" json:"sort"`
	Field string `form:"field" json:"field" binding:"required"`
}

// ListSearchVideoService 视频列表
type ListSearchVideoService struct {
	Number     *int     `form:"number" json:"number"`
	Sorts      []Sorts  `form:"sorts" json:"sorts"`
	SAID       int      `form:"said" json:"said"`
	Searchs    []string `form:"searchs" json:"searchs"`
	PageNumber int      `form:"page_number" json:"page_number"`
	UID        int      `form:"uid" json:"uid"`
}

// List 视频列表服务
func (service *ListSearchVideoService) List() serializer.Response {
	videos := []model.Video{}
	db := model.DB
	//每页的长度 不传入默认20 传入0为不限制长度,此时不需要传入当前页PageNumber
	number := 0
	if service.Number == nil {
		number = 20
		db = db.Limit(number)
	} else {
		log.Println("Number:", *service.Number)
		number = *service.Number
		db = db.Limit(number)
	}
	//添加多个排序条件
	if service.Sorts != nil {
		log.Println("Sorts:", service.Sorts)
		for _, v := range service.Sorts {
			if v.Field != "" {
				//拼接字符串
				strs := []string{v.Field, " ", v.Sort}
				db.Order(util.Join(strs))
			}
		}
	}
	//判断视频状态
	//这里需要先使用一次db.Where(),才能连接后续的db.Or()
	db.Where("state = ?", true)
	//是否分区查找
	if service.SAID != 0 {
		log.Println("SAID:", service.SAID)
		db = db.Where("said = ?", service.SAID)
	}
	//UID查找
	if service.UID != 0 {
		db = db.Where("uid = ?", service.UID)
	}
	//关键字查找
	if service.Searchs != nil {
		log.Println("Searchs:", service.Searchs)
		//循环添加需要查询的关键字
		for _, v := range service.Searchs {
			if v != "" {
				str := util.JoinLike(v)
				//先查找标题
				db.Or("title like ?", str)
				//然后查找昵称
				users := []model.User{}
				model.DB.Where("nickname like ?", str).Find(&users)
				for _, u := range users {
					db.Or("uid = ?", u.ID)
				}
			}
		}
	}
	//分页 当前页数 无参数时默认为第1页
	//db.Offset() 必须放在find()前面，不然查询语句可能不对
	pageNumber := 1
	if service.PageNumber > 1 && number > 0 {
		log.Println("Offset:", number)
		pageNumber = service.PageNumber
		db.Offset((pageNumber - 1) * number)
	}
	//查询
	if err := db.Find(&videos).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "视频查询错误",
			Error: err.Error(),
		}
	}
	//反回数据
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildVideos(videos, pageNumber, number),
		Msg:  "成功",
	}
}
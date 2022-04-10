package user

import (
	"log"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/util"

	"github.com/gin-gonic/gin"
)

// Sorts 排序条件列表
type Sorts struct {
	Sort  string `form:"sort" json:"sort"`
	Field string `form:"field" json:"field" binding:"required"`
}

// 获取用户列表服务
type GetUserListService struct {
	Search     string  `form:"search" json:"search"`
	PageNumber int     `form:"page_number" json:"page_number"`
	Number     *int    `form:"number" json:"number"`
	Sorts      []Sorts `form:"sorts" json:"sorts"`
}

// 用户列表函数
func (service *GetUserListService) List(c *gin.Context) serializer.Response {
	users := []model.User{}
	db := model.DB
	// db.Model(model.User{})
	//总长度
	var total int64
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
	db.Where("root = ?", false)
	//关键字查找
	if service.Search != "" {
		log.Println("Searchs:", service.Search)
		str := util.JoinLike(service.Search)
		//先查找用户名
		db.Where("user_name like ?", str)
		//然后查找昵称
		db.Or("nickname like ?", str)
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
	if err := db.Find(&users).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "用户查询错误",
			Error: err.Error(),
		}
	}
	//查询长度
	db.Limit(-1).Offset(-1)
	if err := db.Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "用户查询错误",
			Error: err.Error(),
		}
	}
	//反回数据
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildUsers(users, pageNumber, number, int(total)),
		Msg:  "成功",
	}
}

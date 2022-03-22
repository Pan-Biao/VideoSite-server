package subArea

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// CreateSubAreaService 创建区域的服务
type CreateSubAreaService struct {
	Name string `form:"name" json:"name" binding:"required,min=2,max=6"`
}

// 视频投稿的服务
func (service *CreateSubAreaService) Create(c *gin.Context) serializer.Response {
	//获取当前用户
	user := model.User{}
	if d, _ := c.Get("user"); d != nil {
		if u, ok := d.(*model.User); ok {
			user = *u
		}
	}
	if !user.Root {
		return serializer.Response{
			Code: 9999,
			Msg:  "没有权限",
		}
	}
	subArea := model.SubArea{
		Name: service.Name,
	}
	if CheckingSubArea(service.Name) {
		return serializer.Response{
			Code: 404,
			Msg:  "名称重复",
		}
	}
	err := model.DB.Create(&subArea).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "分区创建失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildSubArea(subArea),
		Msg:  "成功",
	}
}

func CheckingSubArea(name string) bool {
	subArea := model.SubArea{}
	model.DB.Where("name = ?", name).First(&subArea)
	return subArea.Name != ""
}

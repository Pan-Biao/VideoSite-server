package subArea

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

// UpdateSubAreaService 分区更新服务
type UpdateSubAreaService struct {
	Name string `form:"name" json:"name" binding:"required,min=2,max=6"`
}

// 分区更新
func (service *UpdateSubAreaService) Update(c *gin.Context) serializer.Response {
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
	id := c.Param("id")
	subArea := model.SubArea{}
	err := model.DB.First(&subArea, id).Error

	if CheckingSubArea(service.Name) {
		return serializer.Response{
			Code: 404,
			Msg:  "名称重复",
		}
	}
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "分区不存在",
			Error: err.Error(),
		}
	}
	//更新分区信息
	subArea.Name = service.Name

	err = model.DB.Save(&subArea).Error
	if err != nil {
		return serializer.Response{
			Code:  60002,
			Msg:   "分区更新失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildSubArea(subArea),
		Msg:  "成功",
	}
}

package subArea

import (
	"unicode/utf8"
	"vodeoWeb/model"
	"vodeoWeb/serializer"
	"vodeoWeb/service/funcs"

	"github.com/gin-gonic/gin"
)

// CreateSubAreaService 创建区域的服务
type CreateSubAreaService struct {
	Name string `form:"name" json:"name"`
}

func (service *CreateSubAreaService) Create(c *gin.Context) serializer.Response {
	if utf8.RuneCountInString(service.Name) < 2 || utf8.RuneCountInString(service.Name) > 6 {
		return serializer.ParamErr("名称长度应为2-6")
	}

	//检测权限
	if !funcs.CheckRoot(c) {
		return serializer.CheckNoRight()
	}

	subArea := model.SubArea{
		Name: service.Name,
	}

	if CheckingSubArea(service.Name) {
		return serializer.ParamErr("名称重复")
	}

	if err := model.DB.Create(&subArea).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.ReturnData("成功", serializer.BuildSubArea(subArea))
}

func CheckingSubArea(name string) bool {
	subArea := model.SubArea{}
	model.DB.Where("name = ?", name).First(&subArea)
	return subArea.Name != ""
}

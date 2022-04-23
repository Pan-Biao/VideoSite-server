package carousel

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"

	"github.com/gin-gonic/gin"
)

type ListCarouselService struct{}

type Cb struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
	Cover string `json:"cover"`
}

func (service *ListCarouselService) List(c *gin.Context) serializer.Response {
	carousels := []model.Carousel{}

	if err := model.DB.Find(&carousels).Error; err != nil {
		return serializer.DBErr("", err)
	}

	cb := []Cb{}

	for _, item := range carousels {
		cb = append(cb, Cb{
			ID:    item.ID,
			Title: item.Title,
			Path:  item.Path,
			Cover: item.Cover,
		})
	}

	return serializer.ReturnData("成功", cb)
}

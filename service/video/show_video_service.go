package video

import (
	"vodeoWeb/model"
	"vodeoWeb/serializer"
)

// CreateVideoService 视频详情的服务
type ShowVideoService struct{}

// Show 视频详情
func (service *ShowVideoService) Show(videoId string) serializer.Response {
	video := model.Video{}
	count := int64(0)
	model.DB.First(&video, videoId).Count(&count)
	if count == 0 {
		return serializer.ReturnData("视频不存在", false)
	}

	return serializer.ReturnData("成功", serializer.BuildVideo(video))
}

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
	err := model.DB.First(&video, videoId).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildVideo(video),
		Msg:  "成功",
	}
}

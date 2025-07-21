package video

import (
	"byte_douyin_project/models"
	"byte_douyin_project/util"
)

// PostVideo 投稿视频
func PostVideo(userId int64, videoName, coverName, title string) error {
	videoName = util.GetFileUrl(videoName)
	coverName = util.GetFileUrl(coverName)

	video := &models.Video{
		UserInfoId: userId,
		PlayUrl:    videoName,
		CoverUrl:   coverName,
		Title:      title,
	}
	if err := models.NewVideoDAO().AddVideo(video); err != nil {
		return err
	}
	return nil
}

package comment

import (
	"byte_douyin_project/models"
	"byte_douyin_project/util"
	"errors"
	"fmt"
)

func QueryCommentList(userId, videoId int64) ([]*models.Comment, error) {
	if !models.NewUserInfoDAO().IsUserExistById(userId) {
		return nil, fmt.Errorf("用户%d处于登出状态", userId)
	}
	if !models.NewVideoDAO().IsVideoExistById(videoId) {
		return nil, fmt.Errorf("视频%d不存在或已经被删除", videoId)
	}
	var comments []*models.Comment
	if err := models.NewCommentDAO().QueryCommentListByVideoId(videoId, &comments); err != nil {
		return nil, err
	}
	//根据前端的要求填充正确的时间格式
	if err := util.FillCommentListFields(&comments); err != nil {
		return nil, errors.New("暂时还没有人评论")
	}

	return comments, nil
}

package comment

import (
	"byte_douyin_project/models"
	"byte_douyin_project/util"
	"errors"
	"fmt"
)

const (
	CREATE = 1
	DELETE = 2
)

func PostComment(userId int64, videoId int64, commentId int64, actionType int64, commentText string) (*models.Comment, error) {
	if !models.NewUserInfoDAO().IsUserExistById(userId) {
		return nil, fmt.Errorf("用户%d不存在", userId)
	}
	if !models.NewVideoDAO().IsVideoExistById(videoId) {
		return nil, fmt.Errorf("视频%d不存在", videoId)
	}
	if actionType != CREATE && actionType != DELETE {
		return nil, errors.New("未定义的行为")
	}
	var comment models.Comment
	switch actionType {
	case CREATE:
		comment = models.Comment{UserInfoId: userId, VideoId: videoId, Content: commentText}
		err := models.NewCommentDAO().AddCommentAndUpdateCount(&comment)
		if err != nil {
			return nil, err
		}
	case DELETE:

		if err := models.NewCommentDAO().QueryCommentById(commentId, &comment); err != nil {
			return nil, err
		}
		//删除comment
		if err := models.NewCommentDAO().DeleteCommentAndUpdateCountById(commentId, videoId); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("未定义的操作")
	}
	//填充字段
	userInfo := models.UserInfo{}
	_ = models.NewUserInfoDAO().QueryUserInfoById(comment.UserInfoId, &userInfo)
	comment.User = userInfo
	_ = util.FillCommentFields(&comment)

	return &comment, nil
}

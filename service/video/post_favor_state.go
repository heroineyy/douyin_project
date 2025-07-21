package video

import (
	"byte_douyin_project/cache"
	"byte_douyin_project/models"
	"errors"
)

const (
	PLUS  = 1
	MINUS = 2
)

func PostFavorState(userId, videoId, actionType int64) error {
	if !models.NewUserInfoDAO().IsUserExistById(userId) {
		return errors.New("用户不存在")
	}
	if actionType != PLUS && actionType != MINUS {
		return errors.New("未定义的行为")
	}

	switch actionType {
	case PLUS:
		//视频点赞数目+1
		if err := models.NewVideoDAO().PlusOneFavorByUserIdAndVideoId(userId, videoId); err != nil {
			return errors.New("不要重复点赞")
		}
		//对应的用户是否点赞的映射状态更新
		cache.NewProxyIndexMap().UpdateVideoFavorState(userId, videoId, true)

	case MINUS:
		//视频点赞数目-1
		if err := models.NewVideoDAO().MinusOneFavorByUserIdAndVideoId(userId, videoId); err != nil {
			return errors.New("点赞数目已经为0")
		}
		//对应的用户是否点赞的映射状态更新
		cache.NewProxyIndexMap().UpdateVideoFavorState(userId, videoId, false)

	default:
		return errors.New("未定义的操作")
	}
	return nil
}

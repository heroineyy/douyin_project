package user_info

import (
	"byte_douyin_project/cache"
	"byte_douyin_project/models"
)

func QueryFollowerList(userId int64) ([]*models.UserInfo, error) {
	if !models.NewUserInfoDAO().IsUserExistById(userId) {
		return nil, ErrUserNotExist
	}
	var followerList []*models.UserInfo
	err := models.NewUserInfoDAO().GetFollowerListByUserId(userId, &followerList)
	if err != nil {
		return nil, err
	}
	//填充is_follow字段
	for _, v := range followerList {
		v.IsFollow = cache.NewProxyIndexMap().GetUserRelation(userId, v.Id)
	}
	return followerList, nil
}

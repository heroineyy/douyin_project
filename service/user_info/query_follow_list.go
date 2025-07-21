package user_info

import (
	"byte_douyin_project/models"
	"errors"
)

var (
	ErrUserNotExist = errors.New("用户不存在或已注销")
)

type FollowList struct {
	UserList []*models.UserInfo `json:"user_list"`
}

func QueryFollowList(userId int64) (*FollowList, error) {
	userinfoDao := models.NewUserInfoDAO()
	if !userinfoDao.IsUserExistById(userId) {
		return nil, ErrUserNotExist
	}
	var userList []*models.UserInfo
	err := userinfoDao.GetFollowListByUserId(userId, &userList)
	if err != nil {
		return nil, err
	}
	for i, _ := range userList {
		userList[i].IsFollow = true //当前用户的关注列表，故isFollow定为true
	}
	followList := &FollowList{UserList: userList}

	return followList, nil
}

package video

import (
	"byte_douyin_project/cache"
	"byte_douyin_project/models"
	"errors"
)

func QueryVideoListByUserId(userId int64) (videos []*models.Video, err error) {
	userInfoDAO := models.NewUserInfoDAO()
	if !userInfoDAO.IsUserExistById(userId) {
		return nil, errors.New("用户不存在")
	}
	if err = models.NewVideoDAO().QueryVideoListByUserId(userId, &videos); err != nil {
		return nil, err
	}

	//作者信息查询
	var userInfo models.UserInfo
	if err = userInfoDAO.QueryUserInfoById(userId, &userInfo); err != nil {
		return nil, err
	}
	//填充信息(Author和IsFavorite字段
	for i := range videos {
		videos[i].Author = userInfo
		videos[i].IsFavorite = cache.NewProxyIndexMap().GetVideoFavorState(userId, videos[i].Id)
	}

	return videos, nil
}

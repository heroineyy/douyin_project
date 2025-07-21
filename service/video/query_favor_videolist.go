package video

import (
	"byte_douyin_project/models"
	"errors"
)

func QueryFavorVideoList(userId int64) ([]*models.Video, error) {
	var videos []*models.Video
	if !models.NewUserInfoDAO().IsUserExistById(userId) {
		return nil, errors.New("用户状态异常")
	}
	if err := models.NewVideoDAO().QueryFavorVideoListByUserId(userId, &videos); err != nil {
		return nil, err
	}
	//填充信息(Author和IsFavorite字段，由于是点赞列表，故所有的都是点赞状态
	for i := range videos {
		//作者信息查询
		var userInfo models.UserInfo
		if err := models.NewUserInfoDAO().QueryUserInfoById(videos[i].UserInfoId, &userInfo); err != nil {
			return nil, err
		}
		videos[i].Author = userInfo
		videos[i].IsFavorite = true
	}
	return videos, nil
}

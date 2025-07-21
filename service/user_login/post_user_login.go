package user_login

import (
	"byte_douyin_project/middleware"
	"byte_douyin_project/models"
	"errors"
)

// PostUserLogin 注册用户并得到token和id
func PostUserLogin(username, password string) (*LoginResponse, error) {

	//判断用户名是否已经存在
	userLoginDAO := models.NewUserLoginDao()
	if userLoginDAO.IsUserExistByUsername(username) {
		return nil, errors.New("用户名已存在")
	}

	//更新数据到数据库
	userLogin := models.UserLogin{Username: username, Password: password}
	userinfo := models.UserInfo{User: &userLogin, Name: username}

	//更新操作
	userInfoDAO := models.NewUserInfoDAO()
	err := userInfoDAO.AddUserInfo(&userinfo)
	if err != nil {
		return nil, err
	}

	//颁发token
	token, err := middleware.ReleaseToken(userLogin)
	if err != nil {
		return nil, err
	}
	data := &LoginResponse{
		UserId: userinfo.Id,
		Token:  token,
	}

	return data, nil
}

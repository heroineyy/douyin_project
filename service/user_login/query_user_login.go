package user_login

import (
	"byte_douyin_project/middleware"
	"byte_douyin_project/models"
)

const (
	MaxUsernameLength = 100
	MaxPasswordLength = 20
	MinPasswordLength = 8
)

type LoginResponse struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

// QueryUserLogin 查询用户是否存在，并返回token和id
func QueryUserLogin(username, password string) (*LoginResponse, error) {
	//准备好数据
	userLoginDAO := models.NewUserLoginDao()
	var login models.UserLogin
	//准备好userid
	err := userLoginDAO.QueryUserLogin(username, password, &login)
	if err != nil {
		return nil, err
	}
	//准备颁发token
	token, err := middleware.ReleaseToken(login)
	if err != nil {
		return nil, err
	}
	//打包最终数据
	data := &LoginResponse{
		UserId: login.UserInfoId,
		Token:  token,
	}
	return data, nil
}

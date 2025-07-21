package user_login

import (
	"byte_douyin_project/common"
	"byte_douyin_project/service/user_login"
	"github.com/gin-gonic/gin"
)

func UserLoginHandler(c *gin.Context) {
	username := c.Query("username")
	raw, _ := c.Get("password")
	password, ok := raw.(string)
	if !ok {
		common.ErrorResponse(c, "密码解析错误")
		return
	}

	userLoginResponse, err := user_login.QueryUserLogin(username, password)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}
	common.SuccessResponse(c, userLoginResponse)
}

package user_login

import (
	"byte_douyin_project/common"
	"byte_douyin_project/service/user_login"
	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	rawVal, _ := c.Get("password")
	password, ok := rawVal.(string)
	if !ok {
		common.ErrorResponse(c, "密码解析错误")
		return
	}
	registerResponse, err := user_login.PostUserLogin(username, password)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}
	common.SuccessResponse(c, registerResponse)
}

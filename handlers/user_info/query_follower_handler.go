package user_info

import (
	"byte_douyin_project/common"
	"byte_douyin_project/service/user_info"
	"errors"
	"github.com/gin-gonic/gin"
)

func QueryFollowerHandler(c *gin.Context) {
	userId := c.GetInt64("user_id")
	list, err := user_info.QueryFollowerList(userId)
	if err != nil {
		if errors.Is(err, user_info.ErrUserNotExist) {
			common.ErrorResponse(c, "用户不存在")
		} else {
			common.ErrorResponse(c, "查询失败")
		}
		return
	}

	common.SuccessResponse(c, list)
}

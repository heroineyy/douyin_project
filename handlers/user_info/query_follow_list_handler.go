package user_info

import (
	"byte_douyin_project/common"
	"byte_douyin_project/service/user_info"
	"github.com/gin-gonic/gin"
)

func QueryFollowListHandler(c *gin.Context) {
	userId := c.GetInt64("user_id")

	list, err := user_info.QueryFollowList(userId)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}
	common.SuccessResponse(c, list)
}

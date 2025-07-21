package video

import (
	"byte_douyin_project/common"
	"byte_douyin_project/service/video"
	"github.com/gin-gonic/gin"
)

func QueryVideoListHandler(c *gin.Context) {
	// todo:下面获取参数的代码需要修改
	userId := c.GetInt64("user_id")

	videoList, err := video.QueryVideoListByUserId(userId)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	common.SuccessResponse(c, videoList)
}

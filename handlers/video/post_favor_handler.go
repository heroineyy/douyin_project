package video

import (
	"byte_douyin_project/common"
	"byte_douyin_project/service/video"
	"github.com/gin-gonic/gin"
)

func PostFavorHandler(c *gin.Context) {
	// todo:下面获取参数的代码需要修改
	userId := c.GetInt64("user_id")
	videoId := c.GetInt64("video_id")
	actionType := c.GetInt64("action_type")

	err := video.PostFavorState(userId, videoId, actionType)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	common.SuccessResponse(c, nil)
}

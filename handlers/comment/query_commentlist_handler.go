package comment

import (
	"byte_douyin_project/common"
	"byte_douyin_project/service/comment"
	"github.com/gin-gonic/gin"
)

func QueryCommentListHandler(c *gin.Context) {
	userId := c.GetInt64("user_id")
	videoId := c.GetInt64("video_id")

	commentList, err := comment.QueryCommentList(userId, videoId)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	common.SuccessResponse(c, commentList)
}

package comment

import (
	"byte_douyin_project/common"
	"byte_douyin_project/service/comment"
	"github.com/gin-gonic/gin"
)

// @Summary 发布评论或者删除评论
// @Description 根据用户ID和视频ID发布评论或者删除评论
// @Tags 评论
// @Accept  json
// @Produce  json
// @Param user_id body int64 true "用户ID"
// @Param video_id body int64 true "视频ID"
// @Param action_type body int64 true "操作类型(CREATE/DELETE)"
// @Param comment_text body string false "评论内容(仅创建时使用)"
// @Param comment_id body int64 false "评论ID(仅删除时使用)"
// @Success 200 {object} CommentResponse "评论成功发布"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Router /comment/action/ [post]

func PostCommentHandler(c *gin.Context) {

	userId := c.GetInt64("user_id")
	videoId := c.GetInt64("video_id")
	actionType := c.GetInt64("action_type")
	var commentText string
	var commentId int64

	switch actionType {
	case comment.CREATE:
		commentText = c.Query("comment_text")
	case comment.DELETE:
		commentId = c.GetInt64("comment_id")
	default:
		common.ErrorResponse(c, "actionType解析出错")
		return
	}

	commentRes, err := comment.PostComment(userId, videoId, commentId, actionType, commentText)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}
	common.SuccessResponse(c, commentRes)
}

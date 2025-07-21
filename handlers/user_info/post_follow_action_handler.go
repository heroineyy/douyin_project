package user_info

import (
	"byte_douyin_project/common"
	"byte_douyin_project/service/user_info"
	"errors"
	"github.com/gin-gonic/gin"
)

func PostFollowActionHandler(c *gin.Context) {
	// todo:下面获取参数的代码需要修改
	userId := c.GetInt64("user_id")
	followId := c.GetInt64("to_user_id")
	actionType := c.GetInt("action_type")

	if err := user_info.PostFollowAction(userId, followId, actionType); err != nil {
		if errors.Is(err, user_info.ErrIvdAct) || errors.Is(err, user_info.ErrIvdFolUsr) {
			common.ErrorResponse(c, err.Error())
		} else {
			common.ErrorResponse(c, "请勿重复关注")
		}
		return
	}

	common.SuccessResponse(c, "关注成功")
}

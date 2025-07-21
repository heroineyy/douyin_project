package video

import (
	"byte_douyin_project/common"
	"byte_douyin_project/service/video"
	"github.com/gin-gonic/gin"
)

func QueryFavorVideoListHandler(c *gin.Context) {
	//解析参数
	// todo:下面获取参数的代码需要修改
	userId := c.GetInt64("user_id")
	//正式调用
	favorVideoList, err := video.QueryFavorVideoList(userId)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}
	//成功返回
	common.SuccessResponse(c, favorVideoList)
}

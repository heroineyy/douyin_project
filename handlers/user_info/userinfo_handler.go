package user_info

import (
	"byte_douyin_project/common"
	"byte_douyin_project/models"
	"github.com/gin-gonic/gin"
)

// 获取用户信息
// @Summary 获取用户信息
// @Description 根据用户ID获取用户详细信息
// @Tags 用户相关
// @Accept json
// @Produce json
// @Param user_id query int64 true "用户ID"
// @Param token query string true "用户token"
// @Success 200 {object} common.SuccessResponse "成功获取用户信息"
// @Failure 400 {object} common.ErrorResponse "用户ID不存在或格式错误"
// @Failure 500 {object} common.ErrorResponse "获取用户信息失败"
// @Router /user [get]
func UserInfoHandler(c *gin.Context) {
	// todo：实际上只要传token就行，这里可以优化
	userId := c.GetInt64("user_id")
	// 查询用户信息
	var userInfo models.UserInfo

	if err := models.NewUserInfoDAO().QueryUserInfoById(userId, &userInfo); err != nil {
		common.ErrorResponse(c, "获取用户信息失败: "+err.Error())
		return
	}

	common.SuccessResponse(c, &userInfo)
}

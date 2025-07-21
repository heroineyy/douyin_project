package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type commonResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// 错误响应辅助函数

func ErrorResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, commonResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

// 成功响应辅助函数
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, commonResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		Data:       data,
	})
}

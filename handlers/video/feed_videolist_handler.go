package video

import (
	"byte_douyin_project/common"
	"byte_douyin_project/middleware"
	"byte_douyin_project/service/video"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func FeedVideoListHandler(c *gin.Context) {
	token, ok := c.GetQuery("token")
	//无登录状态
	if !ok {
		rawTimestamp := c.Query("latest_time")
		var latestTime time.Time
		intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
		if err == nil {
			latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
		}
		videoList, err := video.QueryFeedVideoList(0, latestTime)
		if err != nil {
			common.ErrorResponse(c, err.Error())
			return
		}
		common.SuccessResponse(c, videoList)
	}

	//有登录状态
	if claim, ok := middleware.ParseToken(token); ok {
		//token超时
		if time.Now().Unix() > claim.ExpiresAt {
			common.ErrorResponse(c, "token超时")
		}
		rawTimestamp := c.Query("latest_time")
		var latestTime time.Time
		intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
		if err != nil {
			latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
		}
		//调用service层接口
		videoList, err := video.QueryFeedVideoList(claim.UserId, latestTime)
		if err != nil {
			common.ErrorResponse(c, err.Error())
			return
		}
		common.SuccessResponse(c, videoList)

	}

}

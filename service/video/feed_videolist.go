package video

import (
	"byte_douyin_project/models"
	"byte_douyin_project/util"
	"time"
)

// MaxVideoNum 每次最多返回的视频流数量
const (
	MaxVideoNum = 30
)

type FeedVideoList struct {
	Videos   []*models.Video `json:"video_list,omitempty"`
	NextTime int64           `json:"next_time,omitempty"`
}

func QueryFeedVideoList(userId int64, latestTime time.Time) (*FeedVideoList, error) {
	//所有传入的参数不填也应该给他正常处理
	//上层通过把userId置零，表示userId不存在或不需要
	if userId > 0 {
		//这里说明userId是有效的，可以定制性的做一些登录用户的专属视频推荐
	}

	if latestTime.IsZero() {
		latestTime = time.Now()
	}
	var videos []*models.Video

	if err := models.NewVideoDAO().QueryVideoListByLimitAndTime(MaxVideoNum, latestTime, &videos); err != nil {
		return nil, err
	}

	//如果用户为登录状态，则更新该视频是否被该用户点赞的状态
	newlatestTime, _ := util.FillVideoListFields(userId, &videos) //不是致命错误，不返回

	//准备好时间戳
	var nextTime int64
	if newlatestTime != nil {
		nextTime = (*newlatestTime).UnixNano() / 1e6
	} else {
		nextTime = time.Now().Unix() / 1e6

	}

	feedVideo := &FeedVideoList{
		Videos:   videos,
		NextTime: nextTime,
	}
	return feedVideo, nil
}

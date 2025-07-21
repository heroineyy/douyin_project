package video

import (
	"byte_douyin_project/common"
	"byte_douyin_project/service/video"
	"byte_douyin_project/util"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

var (
	videoIndexMap = map[string]struct{}{
		".mp4":  {},
		".avi":  {},
		".wmv":  {},
		".flv":  {},
		".mpeg": {},
		".mov":  {},
	}
	pictureIndexMap = map[string]struct{}{
		".jpg": {},
		".bmp": {},
		".png": {},
		".svg": {},
	}
)

// PublishVideoHandler 发布视频，并截取一帧画面作为封面
func PublishVideoHandler(c *gin.Context) {
	//准备参数
	// todo:下面获取参数的代码需要修改
	userId := c.GetInt64("user_id")
	title := c.PostForm("title")
	form, err := c.MultipartForm()
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	//支持多文件上传
	files := form.File["data"]
	for _, file := range files {
		suffix := filepath.Ext(file.Filename)    //得到后缀
		if _, ok := videoIndexMap[suffix]; !ok { //判断是否为视频格式
			common.ErrorResponse(c, "文件格式错误")
			continue
		}
		name := util.NewFileName(userId) //根据userId得到唯一的文件名
		filename := name + suffix
		savePath := filepath.Join("./static", filename)
		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			common.ErrorResponse(c, err.Error())
			continue
		}
		//截取一帧画面作为封面
		err = util.SaveImageFromVideo(name, true)
		if err != nil {
			common.ErrorResponse(c, err.Error())
			continue
		}
		//数据库持久化
		err := video.PostVideo(userId, filename, name+util.GetDefaultImageSuffix(), title)
		if err != nil {
			common.ErrorResponse(c, err.Error())
			continue
		}
		common.SuccessResponse(c, file.Filename+"上传成功")
	}
}

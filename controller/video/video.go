package video

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os/exec"
	"path/filepath"
	logic "real_my_tiktok/logic/video"
	"real_my_tiktok/pkg/snowflake"
	"strconv"
)

func PublishActionHandler(c *gin.Context) {
	// 把视频存进了static/video下面
	file, err := c.FormFile("data")
	if err != nil {
		zap.L().Error("c.FormFile failed", zap.Error(err))
	}
	VideoDstPath := "./static/video/" + file.Filename

	err = c.SaveUploadedFile(file, VideoDstPath)
	if err != nil {
		zap.L().Error("c.SaveUploadedFile failed", zap.Error(err))
		return
	}

	// 解析参数并存到上下文中
	videoID := snowflake.GenID()
	c.Set("videoID", videoID)
	UserID, isExist := c.Get("userID") // 在解析token里
	if isExist == false {
		zap.L().Error("UserID that parse token get is false, c.Get is failed", zap.Error(err))
	}
	c.Set("UserID", UserID)
	c.Set("PlayURL", VideoDstPath)
	title := c.PostForm("title")
	c.Set("Title", title)
	// 把封面存进static/cover下面
	staticPath := "static/cover"                             // 静态文件保存路径
	coverFileName := strconv.FormatInt(videoID, 10) + ".jpg" // 封面图文件名
	// 构建保存封面图的完整路径
	coverDstPath := filepath.Join(staticPath, coverFileName)
	// 使用 ffmpeg 提取视频的第一帧作为封面图
	cmd := exec.Command("ffmpeg", "-i", VideoDstPath, "-ss", "00:00:00.001", "-vframes", "1", coverDstPath)
	err = cmd.Run()
	if err != nil {
		zap.L().Error("exec.Command failed", zap.Error(err))
		return
	}
	c.Set("CoverURL", coverDstPath)
	// 进入logic层
	err = logic.PublishVideo(c)
	if err != nil {
		zap.L().Error("logic.Publish failed", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "/publish/action/ success",
	})
}

package logic

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	dao "real_my_tiktok/dao/mysql"
)

func PublishVideo(c *gin.Context) (err error) {

	err = dao.PublishVideo(c)
	if err != nil {
		zap.L().Error("dao.PublishVideo failed", zap.Error(err))
		return err
	}
	return

}

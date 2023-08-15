package dao

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"real_my_tiktok/models"
)

func PublishVideo(c *gin.Context) (err error) {
	videoID, isVideoIDExist := c.Get("videoID")
	if isVideoIDExist == false {
		zap.L().Error("idVideoIDExist did not find videoID in gin.Context")
	}
	PlayURL, isPlayURLExist := c.Get("PlayURL")
	if isPlayURLExist == false {
		zap.L().Error("idPlayURLExist did not find PlayURL in gin.Context")
	}
	CoverURL, isCoverURLExist := c.Get("CoverURL")
	if isCoverURLExist == false {
		zap.L().Error("idCoverURLExist did not find CoverURL in gin.Context")
	}
	Title, isTitleExist := c.Get("Title")
	if isTitleExist == false {
		zap.L().Error("isTitleExist did not find Title in gin.Context")
	}

	UserID, isUserIDExist := c.Get("UserID")
	if isUserIDExist == false {
		zap.L().Error("isUserIDExist did not find UserID in gin.Context")
	}
	u := &models.User{ID: UserID.(int64)}
	err = GetUserInfo(u)
	if err != nil {
		zap.L().Error("GetUserInfo is failed", zap.Error(err))
		return err
	}
	author := models.Author{
		ID:              u.ID,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
		Name:            u.Name,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		IsFollow:        u.IsFollow,
	}
	video := &models.Video{
		ID:            videoID.(int),
		FavoriteCount: 0,
		CommentCount:  0,
		PlayUrl:       PlayURL.(string),
		CoverUrl:      CoverURL.(string),
		Title:         Title.(string),
		IsFavorite:    false,
		Author:        author,
	}
	db.Create(&video)
	return
}

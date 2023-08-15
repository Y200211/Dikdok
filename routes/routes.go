package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"real_my_tiktok/controller/user"
	"real_my_tiktok/controller/video"
	"real_my_tiktok/logger"
	JWT "real_my_tiktok/pkg/jwt"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	baseGroup := r.Group("/douyin")
	// 基础接口
	baseGroup.GET("/feed/")
	baseGroup.POST("/user/register/", user.RegisterHandler)
	baseGroup.POST("/user/login/", user.LogInHandler)
	baseGroup.GET("/user/", JWTMiddleWareByParam(), user.GetUserInfo)
	baseGroup.POST("/publish/action/", JWTMiddleWareByReqBody(), video.PublishActionHandler)
	baseGroup.GET("/publish/list/")
	// 互动接口
	baseGroup.POST("/favorite/action/")
	baseGroup.GET("/favorite/list/")
	baseGroup.POST("/comment/action/")
	baseGroup.GET("/comment/list/")
	// 社交接口
	baseGroup.POST("/relation/action/")
	baseGroup.GET("/relation/follow/list/")
	baseGroup.GET("/relation/follower/list/")
	baseGroup.GET("/relation/friend/list/")
	baseGroup.POST("/message/action/")
	baseGroup.GET("/message/chat/")
	return r
}
func JWTMiddleWareByReqBody() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.FormValue("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 404,
				"status_msg":  "token is null",
			})
			c.Abort()
			return
		}

		myclaim, err := JWT.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 404,
				"status_msg":  "token is not right",
			})
			zap.L().Error("JWT.ParseToken failed", zap.Error(err))
		}
		c.Set("userID", myclaim.UserID)
		c.Next()
	}
}
func JWTMiddleWareByParam() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 404,
				"status_msg":  "token is null",
			})
			c.Abort()
			return
		}

		myclaim, err := JWT.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 404,
				"status_msg":  "token is not right",
			})
			zap.L().Error("JWT.ParseToken failed", zap.Error(err))
		}
		c.Set("userID", myclaim.UserID)
		c.Next()
	}
}

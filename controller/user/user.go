package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	logic "real_my_tiktok/logic/user"
	"real_my_tiktok/models"
	"real_my_tiktok/pkg/jwt"
	"real_my_tiktok/pkg/snowflake"
)

func RegisterHandler(c *gin.Context) {
	u := &models.User{}
	//  获取参数和参数校验(获取用户输入的账号和密码，查询他们是否已经注册过了)
	username := c.Query("username")
	password := c.Query("password")
	u.Name = username
	u.Password = password
	u.ID = snowflake.GenID()
	token, err := jwt.GenToken(u.ID, u.Name)
	if err != nil {
		zap.L().Error("RegisterHandler's jwt.GenToken failed, err: ", zap.Error(err))
		return
	}
	if isOnlyUsername := logic.IsOnlyOneUsername(u); !isOnlyUsername {
		zap.L().Warn("username already exists")
		return
	}
	//  业务逻辑操作
	//  放入数据库中

	if err := logic.Register(u); err != nil {
		zap.L().Error("logic.Register failed", zap.Error(err))
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"status_msg": "register success",
		"user_id":    u.ID,
		"token":      token,
	})
}

func LogInHandler(c *gin.Context) {
	u := &models.User{}
	// 获取请求参数 username 和 password
	u.Name = c.Query("username")
	u.Password = c.Query("password")

	// 去数据库查看是否存在该用户，并且密码是否正确

	if err := logic.LogIn(u); err != nil {
		zap.L().Error("logic.LogIn failed", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": 404,
			"status_msg":  err.Error(),
			"user_id":     0,
			"token":       "",
		})
		return
	}
	// 生成 token
	token, err := jwt.GenToken(u.ID, u.Name)
	if err != nil {
		zap.L().Error("jwt.GenToken failed ", zap.Error(err))
	}
	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "login success",
		"user_id":     u.ID,
		"token":       token,
	})
}

func GetUserInfo(c *gin.Context) {

	userID, exist := c.Get("userID")
	if !exist {
		zap.L().Error("GetUserInfo c.Get(:userID is not exist")
	}
	u := &models.User{ID: userID.(int64)}
	err := logic.GetUserInfo(u)
	if err != nil {
		zap.L().Error("logic.GetUserInfo failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "get user Info success",
		"user":        u,
	})
}

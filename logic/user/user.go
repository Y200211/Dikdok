package logic

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	dao "real_my_tiktok/dao/mysql"
	"real_my_tiktok/models"
)

func Register(u *models.User) (err error) {
	if err = dao.InsertUser(u); err != nil {
		fmt.Println("dao.InsertUser failed, err:", err)
		return
	}
	return
}

func IsOnlyOneUsername(u *models.User) bool {
	// 查询数据库，看里面是否已经存在该username
	return dao.IsOnlyUsername(u)
}

func LogIn(u *models.User) (err error) {
	if !iSUserExist(u) { // 如果库里没有该用户
		zap.L().Error("Login: user is not exist")
		err = errors.New("Login: user is not exist")
		return
	}
	// 由于请求参数只有用户名和密码，我们还要拿到userid，因为生成token和返回参数需要userid
	if err := dao.LogIn(u); err != nil {
		zap.Error(err)

		return err
	}
	return
}

func iSUserExist(u *models.User) bool {
	return !dao.IsOnlyUsername(u)
}

func GetUserInfo(u *models.User) (err error) {
	if dao.IsUserExistFromUserID(u) {
		err := dao.GetUserInfo(u)
		if err != nil {
			zap.L().Error("dao.GetUserInfo failed")
			return err
		}
	}
	return
}

package dao

import (
	"errors"
	"real_my_tiktok/models"
)

func InsertUser(u *models.User) (err error) {

	result := db.Create(u)
	return result.Error
}

// 如果数据库表里有u.username的话， return false
func IsOnlyUsername(u *models.User) bool {

	var user models.User

	result := db.First(&user, "username = ?", u.Name)
	if result.Error == nil { // 如果db里面有重复的话
		return false
	} else {
		return true
	}
}

func IsUserExistFromUserID(u *models.User) bool {
	var user models.User
	result := db.First(&user, "user_id = ?", u.ID)
	if result.Error == nil { // 如果db里面有userid的话(userid保证唯一性）
		return true
	} else {
		return false
	}
}

// LogIn 由于请求参数只有用户名和密码，我们还要拿到userid，因为生成token和返回参数需要userid
func LogIn(u *models.User) (err error) {
	// 通过 username 来查询 userid
	var user models.User
	db.First(&user, "username = ?", u.Name)
	if u.Password != user.Password {
		err = errors.New("password failed")
		return
	}
	u.ID = user.ID
	return
}

// 通过userid把数据全读入u中
func GetUserInfo(u *models.User) (err error) {
	result := db.Where("username = ?", u.ID).First(u)
	if result.Error != nil {
		return err
	}
	return
}

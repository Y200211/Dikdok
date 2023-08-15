package dao

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"real_my_tiktok/settings"
)

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(settings.Conf.MysqlLogLevel)),
	})
	if err != nil {
		zap.L().Error("connet mysql failed", zap.Error(err))
	}
	return
}

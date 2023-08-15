package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"real_my_tiktok/dao/mysql"
	"real_my_tiktok/logger"
	"real_my_tiktok/pkg/snowflake"
	"real_my_tiktok/routes"
	"real_my_tiktok/settings"
	"syscall"
	"time"
)

func main() {
	// 1. 加载配置

	if err := settings.Init(); err != nil {
		fmt.Println("settings.Init failed, err: ", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Println("logger.Init failed, err: ", err)
		return
	}
	zap.L().Sync()
	zap.L().Debug("init logger success")
	// 3. 初始化mysql
	if err := dao.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Println("mysql init failed, err: ", err)
		return
	}

	// 初始化雪花算法
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineId); err != nil {
		fmt.Println("snowflake failed, err:", err)
		return
	}
	// 4. 注册路由
	r := routes.Setup()
	// 5. 启动(优雅关机)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: ", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}

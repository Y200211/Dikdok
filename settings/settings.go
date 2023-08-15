package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Port         string `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineId    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
}
type LogConfig struct {
	Level      string `mapstructure:"debug"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Port          int    `mapstructure:"port"`
	MysqlLogLevel int    `mapstructure:"mysql_log_level"`
	Host          string `mapstructure:"host"`
	User          string `mapstructure:"user"`
	Password      string `mapstructure:"password"`
	Dbname        string `mapstructure:"dbname"`
}

func Init() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	err = viper.ReadInConfig() // 查找并读取配置文件

	if err != nil {
		fmt.Println("viper.ReadInConfig failed, err:", err)
	}
	err = viper.Unmarshal(Conf)

	if err != nil {
		fmt.Println("viper.Unmarshal failed, err: ", err)
		return
	}
	if err != nil {
		// 处理读取配置文件的错误
		fmt.Println("viper.ReadInConfig failed, err:", err)
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...", in.Name)
		err = viper.Unmarshal(Conf)
		if err != nil {
			fmt.Println("viper.Unmarshal failed, err: ", err)
			return
		}
	})
	return
}

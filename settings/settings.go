package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name           string `mapstructure:"name"`
	Mode           string `mapstructure:"mode"`
	Port           int    `mapstructure:"port"`
	Version        string `mapstructure:"version"`
	StartTime      string `mapstructure:"start_time"`
	MachineId      int64  `mapstructure:"machine_id"`
	*LogConfig     `mapstructure:"log"`
	*MySQLConfig   `mapstructure:"mysql"`
	*RedisConfig   `mapstructure:"redis"`
	*MongodbConfig `mapstructure:"mongodb"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type MongodbConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	User     string `mapstructure:"user"`
	DB       string `mapstructure:"db"`
	Port     int    `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func Init() (err error) {
	viper.SetConfigFile("./conf/config.yaml") // 指定配置文件
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml") // 专用于远程获取配置信息时：etcd
	//viper.AddConfigPath(".") // 指定查找配置文件的路径

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config files were changed")
		viper.Unmarshal(&Conf)
	})

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err: %v\n", err)
		return
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}
	return
}

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var _cfg MasterConfig

// MasterConfig App主配置
type MasterConfig struct {
	AppDb  map[string]DbConfig
	AppLog LogConfig
}

// DbConfig 数据库配置
type DbConfig struct {
	Host     string
	Port     int
	Db       string
	User     string
	Password string
}

// LogConfig 日志文件配置
type LogConfig struct {
	Path string
	File string
}

// Init 初始化配置文件
func Init() error {
	if err := initConfig(); err != nil {
		panic(err)
	}
	return nil
}

func initConfig() error {
	v := viper.New()
	v.SetConfigName("app")
	v.AddConfigPath("./config")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		fmt.Fprintf(os.Stderr, "%#v", err)
	}

	if err := v.Unmarshal(&_cfg); err != nil {
		panic(err)
	}

	return nil
}

// GetDbConfig 获取数据库配置
func GetDbConfig(db string) DbConfig {
	return _cfg.AppDb[db]
}

// GetLogConfig 获取日志配置
func GetLogConfig() LogConfig {
	return _cfg.AppLog
}

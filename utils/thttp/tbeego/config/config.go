package config

import (
	"github.com/BurntSushi/toml"
)

var _cfg MasterConfig

// MasterConfig App主配置
type MasterConfig struct {
	Db DbConfig  `toml:"database"`
	LC LogConfig `toml:"log"`
}

// DbConfig 数据库配置
type DbConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Db       string `toml:"db"`
	Db1      string `toml:"db1"`
	Db2      string `toml:"db2"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

// LogConfig 日志文件配置
type LogConfig struct {
	Path string `toml:"path"`
	File string `toml:"file"`
}

// Init 初始化配置文件
func Init(cfg string) error {
	if _, err := toml.DecodeFile(cfg, &_cfg); err != nil {
		return err
	}
	return nil
}

// GetDbConfig 获取数据库配置
func GetDbConfig() DbConfig {
	return _cfg.Db
}

// GetLogConfig 获取日志配置
func GetLogConfig() LogConfig {
	return _cfg.LC
}

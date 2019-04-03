package log

import (
	"github.com/astaxie/beego/logs"
	"homework/utils/thttp/tbeego/config"
)

var hlog *logs.BeeLogger

// Init 初始化
func Init() error {
	hlog = logs.NewLogger(100)
	lconf := config.GetLogConfig()
	logFile := lconf.Path + lconf.File
	return hlog.SetLogger(logs.AdapterFile, `{"filename":"`+logFile+`","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":120, "rotate":false}`)
}

// Log 获取日志句柄
func Log() *logs.BeeLogger {
	return hlog
}

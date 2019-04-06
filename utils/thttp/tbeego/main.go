package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"homework/utils/thttp/tbeego/config"
	"homework/utils/thttp/tbeego/log"
	"homework/utils/thttp/tbeego/models/repository/db"
	_ "homework/utils/thttp/tbeego/routers"
	"os"
)

var appConf = flag.String("conf", "app.conf", "server config file")

func init() {
	flag.Parse()
	// 配置文件初始化
	if err := config.Init(*appConf); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// 配置文件初始化
	if err :=
	//初始化日志组件
	if err := log.Init(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// 初始化数据库连接
	if err := db.Init(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func main() {
	l := log.Log()
	l.Info("start server")
	beego.Run()
}

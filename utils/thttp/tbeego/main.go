package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/config/yaml"
	_ "github.com/go-sql-driver/mysql"
	"homework/utils/thttp/tbeego/config"
	"homework/utils/thttp/tbeego/log"
	"homework/utils/thttp/tbeego/models/repository/db"
	_ "homework/utils/thttp/tbeego/routers"
	"os"
)

func init() {
	// BeeGo 配置文件初始化
	if err := beego.LoadAppConfig("yaml", "config/app.yaml"); err != nil {
		os.Exit(-1)
	}
	// 其他配置文件初始化
	if err := config.Init(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// 初始化日志组件
	if err := log.Init(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//// 初始化数据库连接
	if err := db.Init(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func main() {

	beego.Run()
}

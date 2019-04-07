package db

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"homework/utils/thttp/tbeego/config"
	"homework/utils/thttp/tbeego/models"
)

const DB_DNS = "%s:%s@tcp(%s:%d)/%s?charset=utf8"

var _session *dbr.Session

func Init() error {
	//初始化数据连接
	cfg := config.GetDbConfig("db2")
	fmt.Printf("aaa%v", cfg.Password)

	dns := fmt.Sprintf(DB_DNS, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Db)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dns, 30)
	orm.RegisterModel(new(models.User1))
	orm.RegisterModel(new(models.User2))
	//conn, err := dbr.Open("mysql", dns, nil)
	//if err != nil {
	//	return err
	//}
	//
	//_session = conn.NewSession(nil)
	//_session.SetMaxOpenConns(20)
	//_session.SetMaxIdleConns(10)
	//
	//if err = _session.Ping(); err != nil {
	//	return err
	//}
	return nil
}

// GetSession 获取数据库链接会话句柄
func GetSession() *dbr.Session {
	return _session
}

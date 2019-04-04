package db

import (
	"fmt"
	"homework/utils/thttp/tbeego/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

const DB_DNS = "%s:%s@tcp(%s:%d)/%s?loc=Local&parseTime=true"

var _session *dbr.Session

func Init() error {
	//初始化数据连接
	cfg := config.GetDbConfig()

	dns := fmt.Sprintf(DB_DNS, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Db)
	conn, err := dbr.Open("mysql", dns, nil)
	if err != nil {
		return err
	}

	_session = conn.NewSession(nil)
	_session.SetMaxOpenConns(20)
	_session.SetMaxIdleConns(10)

	if err = _session.Ping(); err != nil {
		return err
	}
	return nil
}

// GetSession 获取数据库链接会话句柄
func GetSession() *dbr.Session {
	return _session
}

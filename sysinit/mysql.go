package sysinit

import (
	_ "hello/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func dbinit(aliases  ...string) {
	if len(aliases) <= 0 {
		return
	}
	// 判断开发模式,显示命令行信息
	isDev := ("dev" == beego.AppConfig.String("runmode"))
	if isDev {
		orm.Debug = isDev
	}
	if len(aliases) > 0 {
		for _,alias := range aliases {
			registerDatabase(alias)
			// 主库,自动创建数据表
			if "w" == alias {
				orm.RunSyncdb("default",false,true)
			}
		}
	}else {
		registerDatabase("w")
		orm.RunSyncdb("default",false,true)
	}

}

func registerDatabase(alias string) {
	if len(alias) <= 0 {
		return
	}
	// 连接名称
	dbAlias := alias // default
	if "w" == alias || "default" == alias || len(alias) <= 0 {
		dbAlias = "default"
		alias = "w"
	}
	// 拼接链接uri
	// 数据库:
	dbName := beego.AppConfig.String("db_"+alias+"_database")
	// 数据库用户名:
	dbUser := beego.AppConfig.String("db_"+alias+"_username")
	// 数据库用户密码:
	dbPwd := beego.AppConfig.String("db_"+alias+"_password")
	// 数据库host:
	dbHost := beego.AppConfig.String("db_"+alias+"_host")
	// 数据库port:
	dbPort := beego.AppConfig.String("db_"+alias+"_port")
	// 数据库字符集
	dbCharset := "utf8"
	// root:pwd@tcp(host:port)/dbname?charset=utf8
	orm.RegisterDataBase(dbAlias,"mysql",dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset,30)

}
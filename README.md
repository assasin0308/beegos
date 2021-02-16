# beegos
### 1. 安装

```go
// 官网: 
https://beego.me/
// beego安装:
go get github.com/beego/beego/v2@v2.0.0
// orm安装: 
go get github.com/beego/beego/v2/client/orm

```

### 2. app.conf

```go
appname = hello
httpport = 8080
runmode = dev

#MySQL配置
#主库
db_w_host=127.0.0.1
db_w_port=3306
db_w_username=root
db_w_password=a12345
db_w_database=beegos

#从库
```

### 3. sysinit

```go
// init.go

package sysinit

/**
初始化全局数据链接
 */
func init() {
	dbinit("w")
}
```

```go
// mysql.go

package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func dbinit(alias string) {
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
    
	// 获取配置文件配置
	isDev := ("dev" == beego.AppConfig.String("runmode"))
	if isDev {
		orm.Debug = isDev
	}
    // 自动创建数据表
	if "w" == alias {
		orm.RunSyncdb("default",false,true)
	}

}
```

```go
// sysinit.go

package sysinit

import (
	"github.com/astaxie/beego"
	"path/filepath"
)

func sysinit () {
    // 配置上传文件目录
	uploads := filepath.Join("./","uploads")
	beego.BConfig.WebConfig.StaticDir["/uploads"] = uploads

	// 注册前端使用函数
	registerFunctions()
}

func registerFunctions() {

}

```



### 4. 

```go

```

### 5. 

```go

```

### 6. 

```go

```

### 7. 

```go

```

### 8. 

```go

```

### 9. 

```go

```

### 10. 

```go

```


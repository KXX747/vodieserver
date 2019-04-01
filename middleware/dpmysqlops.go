package middleware

//数据库连接层

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

//mysql的连接层

var(
	engineConn *xorm.Engine
	err  error
	devSql ="root:@tcp(127.0.0.1:3306)/stream?charset=utf8"
)

func init() {

	go ConnMysqlDriver()
	//engineConn.Sync(new(UserTokenAccount))
}

//获取mysql的连接信息
func GetMysqlEngineConn() (*xorm.Engine ,error) {

	return engineConn,err
}


//连接mysql驱动
func ConnMysqlDriver() (*xorm.Engine ,error){

	engine, err := xorm.NewEngine("mysql", devSql)
	if err!=nil{
		panic(err.Error())
	}
	engineConn =engine
	engineConn.ShowSQL(true)
	//设置连接池的空闲数大小
	engineConn.SetMaxIdleConns(100)
	//设置最大打开连接数
	engineConn.SetMaxOpenConns(100)


	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engineConn.SetTableMapper(core.SnakeMapper{})

	return engineConn,err
}






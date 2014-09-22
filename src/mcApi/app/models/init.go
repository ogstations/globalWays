package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"path/filepath"
	"github.com/revel/revel"
	"github.com/revel/config"
)

//数据库连接
var (
	//读数据
	ReaderEngine *xorm.Engine
	//写数据
	WriterEngine *xorm.Engine
)

func init() {
	revel.OnAppStart(initDB)
}

//初始化数据库
func initDB() {

	config_file := (filepath.Join(revel.BasePath, "conf", "databases.conf"))
	c, _ := config.ReadDefault(config_file)

	read_driver, _ := c.String("database", "db.read.driver")
	read_dbName, _ := c.String("database", "db.read.dbname")
	read_user, _ := c.String("database", "db.read.user")
	read_password, _ := c.String("database", "db.read.password")
	read_host, _ := c.String("database", "db.read.host")
	read_encoding, _ := c.String("database", "db.read.encoding")

	//数据库连接
	var err error
	ReaderEngine, err = xorm.NewEngine(read_driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", read_user, read_password, read_host, read_dbName, read_encoding))
	if err != nil {
		revel.WARN.Fatalf("数据库连接错误: %v", err)
	}
	ReaderEngine.SetMapper(core.SameMapper{})

	write_driver, _ := c.String("database", "db.write.driver")
	write_dbname, _ := c.String("database", "db.write.dbname")
	write_user, _ := c.String("database", "db.write.user")
	write_password, _ := c.String("database", "db.write.password")
	write_host, _ := c.String("database", "db.write.host")
	write_encoding, _ := c.String("database", "db.write.encoding")

	WriterEngine, err = xorm.NewEngine(write_driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", write_user, write_password, write_host, write_dbname, write_encoding))
	if err != nil {
		revel.WARN.Fatalf("数据库连接错误: %v",err)
	}
	WriterEngine.SetMapper(core.SameMapper{})

}

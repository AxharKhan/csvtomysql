package main

import (
	_ "csvtomysql/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbConfig := mysql.NewConfig()
	dbConfig.User = "dbAdmin"
	dbConfig.Passwd = "dbPassword"
	dbConfig.Addr = "mysql:3306"
	dbConfig.DBName = "PersonsDb"
	dbConfig.Net = "tcp"

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", dbConfig.FormatDSN())
}

func main() {
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb("default", force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	beego.Run()

}

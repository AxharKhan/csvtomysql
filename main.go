package main

import (
	_ "csvtomysql/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

const dbDriver = "mysql"
const dbName = "PersonsDb"
const dbUserName = "dbAdmin"
const dbPassword = "dbPassword"
const dbTable = "Persons"

func init() {
	orm.RegisterDriver(dbDriver, orm.DRMySQL)

	orm.RegisterDataBase(dbName, dbDriver, dbUserName+":"+dbPassword+"@/"+dbName+"?charset=utf8")
}

func main() {
	beego.Run()

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(dbTable, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

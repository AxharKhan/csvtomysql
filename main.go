package main

import (
	_ "csvtomysql/routers"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func init() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	dbConfig := mysql.NewConfig()
	dbConfig.User = "dbAdmin"
	dbConfig.Passwd = "dbPassword"
	dbConfig.Addr = "mysql:3306"
	dbConfig.DBName = "PersonsDb"
	dbConfig.Net = "tcp"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	flag := false

	// Checks if the Database is ready or not.
	timeoutExceeded := time.After(5 * time.Minute)
	for {
		select {
		case <-timeoutExceeded:
			log.Printf("DB connection failed after %s timeout", 5*time.Minute)
			break

		case <-ticker.C:
			err := orm.RegisterDataBase("default", "mysql", dbConfig.FormatDSN())
			if err == nil {
				flag = true
				log.Println("DB Connection Made Successfully.")

				break
			} else {
				log.Println("DB Connection failed. Retrying...")

			}
			log.Println(errors.Wrapf(err, "Failed to connect to DB. Retrying..."))
		}
		if flag {
			break
		}
	}

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

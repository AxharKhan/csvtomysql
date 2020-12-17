package logic

import (
	"csvtomysql/models"
	"strconv"

	"database/sql"
	"log"

	// uses this driver
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbConfig := mysql.NewConfig()
	dbConfig.User = "dbAdmin"
	dbConfig.Passwd = "dbPassword"
	dbConfig.Addr = "mysql:3306"
	dbConfig.DBName = "PersonsDb"
	dbConfig.Net = "tcp"
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	return db
}

// Create ...Creates the DB Table if not present
func Create() {
	db := dbConn()
	query := "CREATE TABLE IF NOT EXISTS Persons(id int primary key auto_increment, first_name varchar(50), last_name varchar(50), age int, blood_group varchar(50))"
	res, err := db.Exec(query)
	if err != nil {
		panic(err.Error())
	}
	rowsaffected, err := res.RowsAffected()
	t := strconv.FormatInt(rowsaffected, 10)
	log.Printf(t)
	defer db.Close()
}

// Insert ...Inserts an entry into the database
func Insert(person *models.Person) {
	db := dbConn()
	firstname := person.Firstname
	lastname := person.LastName
	age := person.Age
	bloodgroup := person.BloodGroup
	insForm, err := db.Prepare("INSERT INTO Persons(first_name,last_name,age,blood_group) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(firstname, lastname, age, bloodgroup)
	log.Println("INSERT: Name: " + firstname + " | bloodgroup: " + bloodgroup)
	defer db.Close()
}

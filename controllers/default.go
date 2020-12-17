package controllers

import (
	"csvtomysql/models"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
)

const dbDriver = "mysql"
const dbName = "PersonsDb"
const dbUserName = "dbAdmin"
const dbPassword = "dbPassword"
const dbTable = "Persons"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) StoreFiles() {
	c.TplName = "index.html"

	file, header, er := c.GetFile("file") // where <<this>> is the controller and <<file>> the id of your form field
	if er != nil {
		c.Data["FileError"] = "Cannot Read the File please try again."
		return
	}
	if file != nil {
		// get the filename
		fileName := header.Filename

		log.Println("file: " + fileName)

		// Parse the file
		r := csv.NewReader(file)

		persons := []models.Person{}
		count := 0
		var csvdata string = ""
		// Iterate through the records
		for {

			// Read each record from csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				c.Data["FileError"] = "Cannot Read the File please try again."
				return
			}

			o := orm.NewOrm()
			var person models.Person
			var i int64

			person.FirstName = record[0]
			person.LastName = record[1]
			i, err = strconv.ParseInt(record[2], 10, 64)
			if err != nil {
				c.Data["FileError"] = "The data is not in the correct format. Please try another file. \n The correct Format is \"first_name,last_name,age,blood_group\" ."
				return
			}
			person.Age = i
			person.BloodGroup = record[3]

			persons = append(persons, person)
			id, err := o.Insert(&person)

			if err != nil {
				c.Data["FileError"] = "An error occured while trying to store the data, Please try again."
				return
			}
			count++
			csvdata += record[0] + record[1] + record[2] + record[3] + "\n"

			log.Printf("Id: %s Value %s\n", id, record[0], record[1])
		}

		log.Printf(csvdata)

		c.Data["isData"] = true
		c.Data["csvpersons"] = persons

		c.Data["File"] = "file: " + strconv.Itoa(count) + " entries added to the Database."

	}
}

func (c *MainController) ReadDB() {
	c.TplName = "index.html"
	o := orm.NewOrm()

	var persons []*models.Person
	num, err := o.QueryTable("person").All(&persons)

	if err == orm.ErrNoRows {
		c.Data["DBError"] = "There were no entries in the DB."
		return
	} else if err != nil {
		c.Data["DBError"] = "Could not load data from the database."
		return
	}
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	for _, person := range persons {
		fmt.Printf(person.BloodGroup)
	}

	c.Data["dbdata"] = persons
}

func (c *MainController) Post() {

}

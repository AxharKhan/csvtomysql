package controllers

import (
	"csvtomysql/logic"
	"log"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) Post() {
	c.TplName = "index.html"

	file, header, er := c.GetFile("file") // where <<this>> is the controller and <<file>> the id of your form field
	if er != nil {
		panic(er.Error)
	}
	if file != nil {
		// get the filename
		fileName := header.Filename
		// save to server
		err := c.SaveToFile("file", "/"+fileName)
		if err != nil {
			panic(err.Error)
		}
		log.Println("file: " + fileName)
		logic.ReadCSV(fileName)

		c.Data["File"] = "file: " + fileName

	}

}

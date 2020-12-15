package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Person struct {
	Id         int
	Firstname  string
	LastName   string
	Age        int64
	BloodGroup string
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Person))
}

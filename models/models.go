package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Person struct {
	Id         int
	FirstName  string
	LastName   string
	Age        int64
	BloodGroup string
}

func init() {
	orm.RegisterModel(new(Person))

}

package models

import (
	"fmt"
	"log"
	"marketIA/db"

	"github.com/astaxie/beego/orm"
)

type Categorias struct {
	ID     int    `orm:"column(ID)" json:"id"`
	NOMBRE string `orm:"column(NOMBRE)" json:"nombre"`
}

func init() {
	orm.RegisterModel(new(Categorias))
}

func InsertCategories(cat *Categorias) {
	session := db.GetSession()

	if _, err := session.Insert(cat); err != nil {
		log.Println(err)
	}
}
func GetAllCategories() (result []Categorias) {
	session := db.GetSession()

	session.QueryTable("Categorias").All(&result)
	return result
}

func DeleteCategories(id int) {
	session := db.GetSession()
	cat := Categorias{ID: id}
	if num, err := session.Delete(&cat); err != nil {
		log.Println(num)
	}
}

func UpdateCategories(id int, cat Categorias) (err error) {
	session := db.GetSession()
	catFind := Categorias{ID: id}
	if session.Read(&catFind) == nil {
		cat.ID = id
		if num, err := session.Update(&cat); err == nil {
			fmt.Println(num)
		}
	}
	return err
}

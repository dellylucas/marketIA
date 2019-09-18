package models

import (
	"fmt"
	"log"
	"marketIA/db"

	"github.com/astaxie/beego/orm"
)

type Metodos_pago struct {
	ID     int    `orm:"column(ID)" json:"id"`
	NOMBRE string `orm:"column(NOMBRE)" json:"nombre"`
}

func init() {
	orm.RegisterModel(new(Metodos_pago))
}

func InsertPayMethods(payMethods *Metodos_pago) {
	session := db.GetSession()

	if _, err := session.Insert(payMethods); err != nil {
		log.Println(err)
	}
}
func GetAllPayMethods() (result []Metodos_pago) {
	session := db.GetSession()

	session.QueryTable("Metodos_pago").All(&result)
	return result
}

func DeletePayMethods(id int) {
	session := db.GetSession()
	payMethods := Metodos_pago{ID: id}
	if num, err := session.Delete(&payMethods); err != nil {
		log.Println(num)
	}
}

func UpdatePayMethods(id int, payMethods Metodos_pago) (err error) {
	session := db.GetSession()
	payMethodsFind := Metodos_pago{ID: id}
	if session.Read(&payMethodsFind) == nil {
		payMethods.ID = id
		if num, err := session.Update(&payMethods); err == nil {
			fmt.Println(num)
		}
	}
	return err
}

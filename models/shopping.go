package models

import (
	"fmt"
	"log"
	"marketIA/db"

	"github.com/astaxie/beego/orm"
)

type Compras struct {
	ID          int `orm:"column(ID)" json:"id"`
	ID_USUARIO  int `orm:"column(ID_USUARIO)" json:"id_usuario"`
	ID_SERVICIO int `orm:"column(ID_SERVICIO)" json:"id_servicio"`
}

func init() {
	orm.RegisterModel(new(Compras))
}

func InsertShopping(shopping *Compras) {
	session := db.GetSession()

	if _, err := session.Insert(shopping); err != nil {
		log.Println(err)
	}
}
func GetAllShopping() (result []Compras) {
	session := db.GetSession()

	session.QueryTable("Compras").All(&result)
	return result
}

func DeleteShopping(id_user int, id_service int) {
	session := db.GetSession()
	shopping := Compras{ID_USUARIO: id_user, ID_SERVICIO: id_service}
	if num, err := session.Delete(&shopping); err != nil {
		log.Println(num)
	}
}

func UpdateShopping(shopping Compras) (err error) {
	session := db.GetSession()
	shoppingFind := Compras{ID_USUARIO: shopping.ID_USUARIO, ID_SERVICIO: shopping.ID_SERVICIO}
	if session.Read(&shoppingFind) == nil {
		if num, err := session.Update(&shopping); err == nil {
			fmt.Println(num)
		}
	}
	return err
}

package models

import (
	"fmt"
	"log"
	"marketIA/db"

	"github.com/astaxie/beego/orm"
)

type Tienda struct {
	ID           int    `orm:"column(ID)" json:"id"`
	ID_CATEGORIA int    `orm:"column(ID_CATEGORIA)" json:"id_categoria"`
	NOMBRE       string `orm:"column(NOMBRE)" json:"nombre"`
	NIT          string `orm:"column(NIT)" json:"nit"`
	DESCRIPCION  string `orm:"column(DESCRIPCION)" json:"descripcion"`
	CLAVE        string `orm:"column(CLAVE)" json:"clave"`
	CORREO       string `orm:"column(CORREO)" json:"correo"`
}

func init() {
	orm.RegisterModel(new(Tienda))
}

func InsertStore(store *Tienda) {
	session := db.GetSession()

	if _, err := session.Insert(store); err != nil {
		log.Println(err)
	}
}
func GetAllStore() (result []Tienda) {
	session := db.GetSession()
	var store []Tienda

	session.QueryTable("Tienda").All(&store)
	return store
}

func DeleteStore(id int) {
	session := db.GetSession()
	store := Tienda{ID: id}
	if num, err := session.Delete(&store); err != nil {
		log.Println(num)
	}
}

func UpdateStore(id int, store Tienda) (err error) {
	session := db.GetSession()
	storeFind := Tienda{ID: id}
	if session.Read(&storeFind) == nil {
		store.ID = id
		if num, err := session.Update(&store); err == nil {
			fmt.Println(num)
		}
	}
	return err
}

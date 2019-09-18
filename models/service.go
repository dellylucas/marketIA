package models

import (
	"fmt"
	"log"
	"marketIA/db"

	"github.com/astaxie/beego/orm"
)

type Servicio struct {
	ID          int    `orm:"column(ID)" json:"id"`
	ID_TIENDA   int    `orm:"column(ID_TIENDA)" json:"id_tienda"`
	NOMBRE      string `orm:"column(NOMBRE)" json:"nombre"`
	PRECIO      int    `orm:"column(PRECIO)" json:"precio"`
	DESCRIPCION string `orm:"column(DESCRIPCION)" json:"descripcion"`
}

func init() {
	orm.RegisterModel(new(Servicio))
}

func InsertService(service *Servicio) {
	session := db.GetSession()

	if _, err := session.Insert(service); err != nil {
		log.Println(err)
	}
}
func GetAllService() (result []Servicio) {
	session := db.GetSession()
	var service []Servicio

	session.QueryTable("Servicio").All(&service)
	return service
}

func DeleteService(id int) {
	session := db.GetSession()
	service := Servicio{ID: id}
	if num, err := session.Delete(&service); err != nil {
		log.Println(num)
	}
}

func UpdateService(id int, service Servicio) (err error) {
	session := db.GetSession()
	serviceFind := Servicio{ID: id}
	if session.Read(&serviceFind) == nil {
		service.ID = id
		if num, err := session.Update(&service); err == nil {
			fmt.Println(num)
		}
	}
	return err
}

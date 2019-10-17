package models

import (
	"fmt"
	"log"
	"marketIA/db"

	"github.com/astaxie/beego/orm"
)

//Usuarios modelo de usuarios
type Usuarios struct {
	ID        int    `orm:"column(ID)" json:"id"`
	NOMBRE    string `orm:"column(NOMBRE)" json:"nombre"`
	APELLIDO  string `orm:"column(APELLIDO)" json:"apellido"`
	CELULAR   string `orm:"column(CELULAR)" json:"celular"`
	DOCUMENTO string `orm:"column(DOCUMENTO)" json:"documento"`
	CLAVE     string `orm:"column(CLAVE)" json:"clave"`
	CORREO    string `orm:"column(CORREO)" json:"correo"`
	ISADMIN   bool   `orm:"column(IS_ADMIN)" json:"admin"`
}

func init() {
	orm.RegisterModel(new(Usuarios))
}

func InsertUser(user *Usuarios) (err error) {
	session := db.GetSession()

	_, err = session.Insert(user)
	return err
}
func GetAllUser() (result []Usuarios) {
	session := db.GetSession()
	var user []Usuarios

	session.QueryTable("Usuarios").All(&user)
	return user
}

func DeleteUser(id int) {
	session := db.GetSession()
	user := Usuarios{ID: id}
	if num, err := session.Delete(&user); err != nil {
		log.Println(num)
	}
}

func UpdateUser(id int, user Usuarios) (err error) {
	session := db.GetSession()
	userFind := Usuarios{ID: id}
	if session.Read(&userFind) == nil {
		user.ID = id
		if num, err := session.Update(&user); err == nil {
			fmt.Println(num)
		}
	}
	return err
}

func ValidateLogin(user Usuarios) (pos Usuarios) {
	session := db.GetSession()
	session.QueryTable("Usuarios").Filter("CLAVE", user.CLAVE).Filter("CORREO", user.CORREO).One(&pos)

	return pos
}

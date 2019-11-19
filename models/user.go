package models

import (
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
	IMAGEN    string `orm:"column(IMAGEN)" json:"imagen"`
}

func init() {
	orm.RegisterModel(new(Usuarios))
}

func InsertUser(user *Usuarios) (err error) {
	session := db.GetSession()

	_, err = session.Insert(user)
	return err
}

func GetUser(id int) (user Usuarios) {
	session := db.GetSession()

	session.QueryTable("Usuarios").Filter("DOCUMENTO", id).One(&user)
	return user
}

func GetAllUser() (result []Usuarios) {
	session := db.GetSession()
	var user []Usuarios

	session.QueryTable("Usuarios").All(&user)
	return user
}

func DeleteUser(id string) (err error) {
	session := db.GetSession()
	var user Usuarios
	if err = session.QueryTable("Usuarios").Filter("DOCUMENTO", id).One(&user); err == nil {
		_, err = session.Delete(&user)
	}
	return err
}

func UpdateUser(id string, user Usuarios) (err error) {
	session := db.GetSession()
	var userFind Usuarios
	if err = session.QueryTable("Usuarios").Filter("DOCUMENTO", id).One(&userFind); err == nil {
		userFind.NOMBRE = user.NOMBRE
		userFind.APELLIDO = user.APELLIDO
		userFind.CELULAR = user.CELULAR
		userFind.CLAVE = user.CLAVE
		userFind.CORREO = user.CORREO
		userFind.ISADMIN = user.ISADMIN
		_, err = session.Update(&userFind)
	}
	return err
}

func ValidateLogin(user Usuarios) (pos Usuarios) {
	session := db.GetSession()
	session.QueryTable("Usuarios").Filter("CLAVE", user.CLAVE).Filter("CORREO", user.CORREO).One(&pos)

	return pos
}

package controllers

import (
	"encoding/json"
	"marketIA/models"
	"marketIA/utils"
	"strconv"

	"github.com/astaxie/beego"
)

//operaciones sobre usuarios
type UserController struct {
	beego.Controller
}

// @Title Get
// @Description Obtiene todos los usuarios
// @Success 200  models.Usuarios
// @router / [get]
func (o *UserController) Get() {
	ob := models.GetAllUser()

	o.Data[utils.TypeMessage] = ob

	o.ServeJSON()
}

//Post - insert
// @Title Post
// @Description crea o inserta usuario
// @Param	body		body 	models.Usuarios	true		"The object content"
// @Success 200 models.Usuarios
// @Failure 403 body is empty
// @router / [post]
func (o *UserController) Post() {
	var ob models.Usuarios
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	err := models.InsertUser(&ob)
	if err != nil {
		o.Data[utils.TypeMessage] = err
	} else {
		o.Data[utils.TypeMessage] = utils.MessageOK
	}
	o.ServeJSON()
}

// @Title Update
// @Description actualiza usuario
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Usuarios	true		"The body"
// @Success 200 {object} models.Usuarios
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *UserController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Usuarios
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	id, _ := strconv.Atoi(objectId)
	err := models.UpdateUser(id, ob)
	if err != nil {
		o.Data[utils.TypeMessage] = err.Error()
	} else {
		o.Data[utils.TypeMessage] = utils.MessageOK
	}
	o.ServeJSON()
}

// @Title Delete
// @Description Elimina usuario
// @Param	userID		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 userID is empty
// @router /:userID [Delete]
func (o *UserController) Delete() {
	userID := o.Ctx.Input.Param(":userID")
	id, _ := strconv.Atoi(userID)
	models.DeleteUser(id)
	o.Data[utils.TypeMessage] = utils.MessageOK
	o.ServeJSON()
}

//Login - login
// @Title Post
// @Description realiza la validacion del ingreso de un usuario o tienda
// @Param	body		body 	models.Usuarios	true		"Contiene el usuario y contraseña digitado"
// @Success 200 {Boolean} models.Usuarios o false
// @Failure 403 body is empty
// @router /login [post]
func (o *UserController) Login() {
	var user models.Usuarios
	json.Unmarshal(o.Ctx.Input.RequestBody, &user)
	user = models.ValidateLogin(user)
	if user.CORREO != "" {
		o.Data[utils.TypeMessage] = user
	} else {
		o.Data[utils.TypeMessage] = false
	}
	o.ServeJSON()
}

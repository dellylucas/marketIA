package controllers

import (
	"encoding/json"
	"marketIA/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about object
type UserController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Success 200  models.User
// @router / [get]
func (o *UserController) Get() {
	ob := models.GetAllUser()

	o.Data["json"] = ob

	o.ServeJSON()
}

//Post - insert
// @Title Post
// @Description create Tutor
// @Param	body		body 	models.Tutor	true		"The object content"
// @Success 200 models.Tutor
// @Failure 403 body is empty
// @router / [post]
func (o *UserController) Post() {
	var ob models.Usuarios
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	err := models.InsertUser(&ob)
	if err != nil {
		o.Data["json"] = err
	} else {
		o.Data["json"] = "ok"
	}
	o.ServeJSON()
}

// @Title Update
// @Description update the object
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
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	userID		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 userID is empty
// @router /:userID [Delete]
func (o *UserController) Delete() {
	userID := o.Ctx.Input.Param(":userID")
	id, _ := strconv.Atoi(userID)
	models.DeleteUser(id)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

//Login - login
// @Title Post
// @Description login
// @Param	body		body 	models.Usuarios	true		"The object content"
// @Success 200 models.Usuarios
// @Failure 403 body is empty
// @router /login/ [post]
func (o *UserController) Login() {
	var user models.Usuarios
	json.Unmarshal(o.Ctx.Input.RequestBody, &user)
	err := models.ValidateLogin(user)
	if err != nil {
		o.Data["json"] = false
	} else {
		o.Data["json"] = true
	}
	o.ServeJSON()
}

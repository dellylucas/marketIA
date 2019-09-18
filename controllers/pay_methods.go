package controllers

import (
	"encoding/json"
	"marketIA/models"
	"strconv"

	"github.com/astaxie/beego"
)

//Operations about object
type PayMethodsController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Success 200  models.Metodos_pago
// @router / [get]
func (o *PayMethodsController) Get() {
	ob := models.GetAllPayMethods()

	o.Data["json"] = ob

	o.ServeJSON()
}

//Post - insert
// @Title Post
// @Description create Metodos_pago
// @Param	body		body 	models.Metodos_pago	true		"The object content"
// @Success 200 models.Metodos_pago
// @Failure 403 body is empty
// @router / [post]
func (o *PayMethodsController) Post() {
	var ob models.Metodos_pago
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	models.InsertPayMethods(&ob)
	o.Data["json"] = "insert success!"
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Metodos_pago	true		"The body"
// @Success 200 {object} models.Metodos_pago
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *PayMethodsController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Metodos_pago
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	id, _ := strconv.Atoi(objectId)
	err := models.UpdatePayMethods(id, ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	payID		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 payID is empty
// @router /:payID [Delete]
func (o *PayMethodsController) Delete() {
	payID := o.Ctx.Input.Param(":payID")
	id, _ := strconv.Atoi(payID)
	models.DeletePayMethods(id)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

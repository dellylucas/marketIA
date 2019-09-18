package controllers

import (
	"encoding/json"
	"marketIA/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about object
type ServiceController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Success 200  models.Servicio
// @router / [get]
func (o *ServiceController) Get() {
	ob := models.GetAllService()

	o.Data["json"] = ob

	o.ServeJSON()
}

//Post - insert
// @Title Post
// @Description create Servicio
// @Param	body		body 	models.Servicio	true		"The object content"
// @Success 200 models.Servicio
// @Failure 403 body is empty
// @router / [post]
func (o *ServiceController) Post() {
	var ob models.Servicio
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	models.InsertService(&ob)
	o.Data["json"] = "insert success!"
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Servicio	true		"The body"
// @Success 200 {object} models.Servicio
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *ServiceController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Servicio
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	id, _ := strconv.Atoi(objectId)
	err := models.UpdateService(id, ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	serviceID		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 serviceID is empty
// @router /:serviceID [Delete]
func (o *ServiceController) Delete() {
	serviceID := o.Ctx.Input.Param(":serviceID")
	id, _ := strconv.Atoi(serviceID)
	models.DeleteService(id)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

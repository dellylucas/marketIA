package controllers

import (
	"encoding/json"
	"marketIA/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about object
type StoreController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Success 200  models.Tienda
// @router / [get]
func (o *StoreController) Get() {
	ob := models.GetAllStore()

	o.Data["json"] = ob

	o.ServeJSON()
}

//Post - insert
// @Title Post
// @Description create Tienda
// @Param	body		body 	models.Tienda	true		"The object content"
// @Success 200 models.Tienda
// @Failure 403 body is empty
// @router / [post]
func (o *StoreController) Post() {
	var ob models.Tienda
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	models.InsertStore(&ob)
	o.Data["json"] = "insert success!"
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Tienda	true		"The body"
// @Success 200 {object} models.Tienda
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *StoreController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Tienda
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	id, _ := strconv.Atoi(objectId)
	err := models.UpdateStore(id, ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	storeID		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 storeID is empty
// @router /:storeID [Delete]
func (o *StoreController) Delete() {
	storeID := o.Ctx.Input.Param(":storeID")
	id, _ := strconv.Atoi(storeID)
	models.DeleteStore(id)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

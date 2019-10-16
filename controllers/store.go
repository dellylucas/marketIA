package controllers

import (
	"encoding/json"
	"marketIA/models"
	"marketIA/utils"
	"strconv"

	"github.com/astaxie/beego"
)

// operaciones sobre tiendas
type StoreController struct {
	beego.Controller
}

// @Title Get
// @Description obtiene todas las tiendas
// @Success 200  models.Tienda
// @router / [get]
func (o *StoreController) Get() {
	ob := models.GetAllStore()

	o.Data[utils.TypeMessage] = ob

	o.ServeJSON()
}

//Post - insert
// @Title Post
// @Description crea una Tienda
// @Param	body		body 	models.Tienda	true		"The object content"
// @Success 200 models.Tienda
// @Failure 403 body is empty
// @router / [post]
func (o *StoreController) Post() {
	var ob models.Tienda
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	models.InsertStore(&ob)
	o.Data[utils.TypeMessage] = utils.MessageOK
	o.ServeJSON()
}

// @Title Update
// @Description actualiza una tienda
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
		o.Data[utils.TypeMessage] = err.Error()
	} else {
		o.Data[utils.TypeMessage] = utils.MessageOK
	}
	o.ServeJSON()
}

// @Title Delete
// @Description elimina una tienda
// @Param	storeID		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 storeID is empty
// @router /:storeID [Delete]
func (o *StoreController) Delete() {
	storeID := o.Ctx.Input.Param(":storeID")
	id, _ := strconv.Atoi(storeID)
	models.DeleteStore(id)
	o.Data[utils.TypeMessage] = utils.MessageOK
	o.ServeJSON()
}

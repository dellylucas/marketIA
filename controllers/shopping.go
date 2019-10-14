package controllers

import (
	"encoding/json"
	"marketIA/models"
	"marketIA/utils"
	"strconv"

	"github.com/astaxie/beego"
)

// operaciones sobre compras
type ShoppingController struct {
	beego.Controller
}

// @Title Get
// @Description obtiene compras
// @Success 200  models.Compras
// @router / [get]
func (o *ShoppingController) Get() {
	ob := models.GetAllShopping()

	o.Data[utils.TypeMessage] = ob

	o.ServeJSON()
}

//Post - insert
// @Title Post
// @Description crea compras
// @Param	body		body 	models.Compras	true		"The object content"
// @Success 200 models.Compras
// @Failure 403 body is empty
// @router / [post]
func (o *ShoppingController) Post() {
	var ob models.Compras
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	models.InsertShopping(&ob)
	o.Data[utils.TypeMessage] = utils.MessageOK
	o.ServeJSON()
}

// @Title Update
// @Description actualiza compras
// @Param	body		body 	models.Compras	true		"The body"
// @Success 200  models.Compras
// @Failure 403 body is empty
// @router / [put]
func (o *ShoppingController) Put() {
	var ob models.Compras
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	err := models.UpdateShopping(ob)
	if err != nil {
		o.Data[utils.TypeMessage] = err.Error()
	} else {
		o.Data[utils.TypeMessage] = utils.MessageOK
	}
	o.ServeJSON()
}

// @Title Delete
// @Description elimina las compras
// @Param	userID		path 	string	true		"The userID you want to delete"
// @Param	serviceID		path 	string	true		"The serviceID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 serviceID is empty
// @router /:serviceID [Delete]
func (o *ShoppingController) Delete() {
	userID := o.Ctx.Input.Param(":userID")
	serviceID := o.Ctx.Input.Param(":serviceID")
	idUser, _ := strconv.Atoi(userID)
	idService, _ := strconv.Atoi(serviceID)
	models.DeleteShopping(idUser, idService)
	o.Data[utils.TypeMessage] = utils.MessageOK
	o.ServeJSON()
}

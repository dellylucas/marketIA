package controllers

import (
	"encoding/json"
	"marketIA/models"
	"marketIA/utils"
	"strconv"

	"github.com/astaxie/beego"
)

//operaciones sobre las categorias de las tiendas de IA
type CategoriesController struct {
	beego.Controller
}

// @Title Get
// @Description obtiene todas las categorias
// @Success 200  models.Categorias
// @router / [get]
func (o *CategoriesController) Get() {
	ob := models.GetAllCategories()

	o.Data[utils.TypeMessage] = ob

	o.ServeJSON()
}

//Post - insert
// @Title Post
// @Description crea una categoria
// @Param	body		body 	models.Categorias	true		"The object content"
// @Success 200 models.Categorias
// @Failure 403 body is empty
// @router / [post]
func (o *CategoriesController) Post() {
	var ob models.Categorias
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	models.InsertCategories(&ob)
	o.Data[utils.TypeMessage] = utils.MessageOK
	o.ServeJSON()
}

// @Title Update
// @Description actualiza una categoria
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Categorias	true		"The body"
// @Success 200 {object} models.Categorias
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *CategoriesController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Categorias
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	id, _ := strconv.Atoi(objectId)
	err := models.UpdateCategories(id, ob)
	if err != nil {
		o.Data[utils.TypeMessage] = err.Error()
	} else {
		o.Data[utils.TypeMessage] = utils.MessageOK
	}
	o.ServeJSON()
}

// @Title Delete
// @Description elimina una categoria
// @Param	catID		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 catID is empty
// @router /:catID [Delete]
func (o *CategoriesController) Delete() {
	catID := o.Ctx.Input.Param(":catID")
	id, _ := strconv.Atoi(catID)
	models.DeleteCategories(id)
	o.Data[utils.TypeMessage] = utils.MessageOK
	o.ServeJSON()
}

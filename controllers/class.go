package controllers

import (
	"encoding/json"
	"scholy/models"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type ClassController struct {
	beego.Controller
}

// @Title Create
// @router / [post]
func (o *ClassController) Post() {  //ko can xu ly id, chi body
	var class models.Class
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &class)			//covert info tu body trong url sang json

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if class.Name == " " || class.CourseID == 0 {			//check valid student (name va dob)
		o.Data["json"] = "nhap lai di. ngu qua"
	} else {
		st, err := models.AddOneClass(class)

		if err != nil { //check loi db coi ne
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = st
		}
	}
	o.ServeJSON()
}

// @Title Get
// @Description find object by studentId
// @Param	objectId		path 	string	true		"the studentId you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ClassController) Get() {
	var IdClass int64
	o.Ctx.Input.Bind(&IdClass, "id_class")			//convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdClass > 0 {
		st, err := models.GetClassByID(IdClass)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = st
		}
	} else {
		o.Data["json"] = "nhap lai di. ngu qua"
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ClassController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [put]
func (o *ClassController) Put() {
	var class models.Class
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &class) //covert info tu body trong url sang json

	var IdClass int64
	o.Ctx.Input.Bind(&IdClass, "id_class") 			//convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdClass > 0 {
		if class.Name == "" || class.CourseID == 0 {
			o.Data["json"] = "nhap lai di. ngu qua"
		} else {
			st, err := models.UpdateOneClass(IdClass, class) //
			if err != nil {
				o.Data["json"] = err.Error()
			} else {
				o.Data["json"] = st
			}
		}
	} else {
		o.Data["json"] = "nhap lai di. ngu qua"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router / [delete]
func (o *ClassController) Delete() {
	var IdClass int64
	o.Ctx.Input.Bind(&IdClass, "id_class")			//convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdClass > 0 {
		st, err := models.DeleteOneClass(IdClass)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = st
		}
	} else {
		o.Data["json"] = "nhap lai di. ngu qua"
	}



	o.ServeJSON()
}


package controllers

import (
	"encoding/json"
	"fmt"
	"scholy/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type StudentController struct {
	beego.Controller
}

// @Title Create
// @router / [post]
func (o *StudentController) Post() {  //ko can xu ly id, chi body
	var student models.Student
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &student)

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if student.Name == "" || student.DOB == "" {			//check valid student (name vaf dob)
		o.Data["json"] = "nhap lai di. ngu qua"
	} else {
		st, err := models.AddOneStudent(student)

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
func (o *StudentController) Get() {
	var IdStudent int64
	o.Ctx.Input.Bind(&IdStudent, "id_student")

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdStudent > 0 {
		st, err := models.GetStudentByID(IdStudent)
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
func (o *StudentController) GetAll() {
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
func (o *StudentController) Put() {
	var student models.Student
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &student) //covert info tu body trong url sang json

	var IdStudent int64
	o.Ctx.Input.Bind(&IdStudent, "id_student") //convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdStudent > 0 {
		if student.Name == "" || student.DOB == "" {
			o.Data["json"] = "nhap lai di. ngu qua"
		} else {
			st, err := models.UpdateOneStudent(IdStudent, student) //
			if err != nil {
				o.Data["json"] = err.Error()
			} else {
				o.Data["json"] = st
			}
		}
	} else {
			st, err := models.GetAllStudents()
			if err != nil {
				o.Data["json"] = err.Error()
			} else {
				o.Data["json"] = st
			}
		}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router / [delete]
func (o *StudentController) Delete() {
	var IdStudent int64
	o.Ctx.Input.Bind(&IdStudent, "id_student")			//convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdStudent > 0 {
		st, err := models.DeleteOneStudent(IdStudent)
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


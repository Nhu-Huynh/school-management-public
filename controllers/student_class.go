package controllers

import (
	"encoding/json"
	"scholy/models"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type StudentClassController struct {
	beego.Controller
}

// @Title Create
// @router / [post]
func (o *StudentClassController) Post() {  //ko can xu ly id, chi body
	var studentclass models.StudentClass
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &studentclass)			//covert info tu body trong url sang json

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if studentclass.StudentID == 0 || studentclass.ClassID == 0 {			//check valid student (name va dob)
		o.Data["json"] = "nhap lai di. ngu qua"
	} else {
		st, err := models.AddOneStudentClass(studentclass)

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
// @router /by_student [get]
func (o *StudentClassController) GetByStudent() {
	var IdStudent int64                        //convert id trong url
	o.Ctx.Input.Bind(&IdStudent, "id_student") //convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdStudent > 0 {
		st, err := models.GetClassByStudent(IdStudent)
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


// @Title Get
// @Description find object by studentId
// @Param	objectId		path 	string	true		"the studentId you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /by_class [get]
func (o *StudentClassController) GetByClass() {
	var IdClass int64                      //convert id trong url
	o.Ctx.Input.Bind(&IdClass, "id_class") //convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdClass > 0 {
		st, err := models.GetStudentByClass(IdClass)
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

// @Title Get
// @Description find object by studentId
// @Param	objectId		path 	string	true		"the studentId you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /by_studentclass [get]
func (o *StudentClassController) GetStudentClass() {
	var IdStudentClass int64                             //convert id trong url
	o.Ctx.Input.Bind(&IdStudentClass, "id_studentclass") //convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdStudentClass > 0 {
		st, err := models.GetStudentClassByID(IdStudentClass)
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
func (o *StudentClassController) GetAll() {
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
func (o *StudentClassController) Put() {
	var studentclass models.StudentClass
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &studentclass) //covert info tu body trong url sang json

	var IdStudentClass int64
	o.Ctx.Input.Bind(&IdStudentClass, "id_studentclass") 			//convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdStudentClass > 0 {
		if studentclass.StudentID == 0 || studentclass.ClassID == 0 {
			o.Data["json"] = "nhap lai di. ngu qua"
		} else {
			st, err := models.UpdateOneStudentClass(IdStudentClass, studentclass) //
			if err != nil {
				o.Data["json"] = err.Error()
			} else {
				o.Data["json"] = st
			}
		}
	} else {
		st, err := models.GetAllStudentClasses()
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
func (o *StudentClassController) Delete() {
	var IdStudentClass int64
	o.Ctx.Input.Bind(&IdStudentClass, "id_studentclass")			//convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdStudentClass > 0 {
		st, err := models.DeleteOneStudentClass(IdStudentClass)
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


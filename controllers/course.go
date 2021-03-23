package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"scholy/models"
	"fmt"
)

// Operations about object
type CourseController struct {
	beego.Controller
}

// @Title Create
// @router / [post]
func (o *CourseController) Post() {  //ko can xu ly id, chi body
	var course models.Course
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &course)			//covert info tu body trong url sang json

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if course.Name == ""  {			//check valid course
		o.Data["json"] = "nhap lai di. ngu qua"
	} else {
		st, err := models.AddOneCourse(course)

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
func (o *CourseController) Get() {
	var IdCourse int64
	o.Ctx.Input.Bind(&IdCourse, "id_course")			//convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdCourse > 0 {
		st, err := models.GetCourseByID(IdCourse)
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
func (o *CourseController) GetAll() {
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
func (o *CourseController) Put() {
	var course models.Course
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &course) //covert info tu body trong url sang json

	var IdCourse int64
	o.Ctx.Input.Bind(&IdCourse, "id_course") 			//convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdCourse > 0 {
		if course.Name == "" {
			o.Data["json"] = "nhap lai di. ngu qua"
		} else {
			st, err := models.UpdateOneCourse (IdCourse, course) //
			if err != nil {
				o.Data["json"] = err.Error()
			} else {
				o.Data["json"] = st
			}
		}
	} else {
		st, err := models.GetAllCourses()
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
func (o *CourseController) Delete() {
	var IdCourse int64
	o.Ctx.Input.Bind(&IdCourse, "id_course")			//convert id trong url

	a := o.Ctx.Request.Header.Get("Authorization")
	fmt.Println(a)

	if IdCourse > 0 {
		st, err := models.DeleteOneCourse(IdCourse)
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


package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Course struct {
	ID int64				`json:"id"`
	Name string				`json:"name"`
	Value int64				`json:"value"`
	Currency string			`json:"currency"`
	BeginAt *string			`json:"begin_at"`
	EndAt *string			`json:"end_at"`
	//pointer mac dinh la null. string mac dinh la chuoi rong
	Duration int64
	Deleted bool 
}


//ten struct
func (Course) TableName() string {
	return "scholl.courses"
}
func init() {
	dsn := "host=localhost user=postgres password=abc123 dbname=myDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	var course Course
	db.First(&course)
	fmt.Println(course)
}


func GetAllCourses () (courses []Course, err error) {
	db, err := setDatabase();
	if err == nil {
		err = db.Where("deleted = ?", false).Find(&courses).Error
	}
	return
}

func GetCourseByID (IdCourse int64) (course Course, err error) {
	db, err := setDatabase()
	if err == nil {
		err = db.Where("id = ? AND deleted = ?", IdCourse, false).Find(&course).Error
	}
	return
}

func AddOneCourse (course Course) (Course, error) {
	db, err := setDatabase()
	if err == nil {
		err = db.Create(&course).Error
	}
	return course, err
}

func UpdateOneCourse (IDCourse int64, newCourse Course) (Course, error) {
	db, err := setDatabase()
	if err == nil {
		oldCourse, err := GetCourseByID(IDCourse)
		if err == nil && oldCourse.ID != 0 {
			if oldCourse.Name != "" {
				oldCourse.Name = newCourse.Name
			}
			if oldCourse.Value != 0 {
				oldCourse.Value = newCourse.Value
			}
			err = db.Save(oldCourse).Error
			newCourse = oldCourse
		}
	} else {
		fmt.Println("no data :<")
	}
	return newCourse, err
}

func DeleteOneCourse (IDCourse int64) (Course, error) {
	db, err := setDatabase()
	var course Course
	if err == nil {
		course, err = GetCourseByID(IDCourse)
		if err == nil {
			course.Deleted = true
			db.Save(&course)
		}
	} else {
		fmt.Println("no data :<")
	}
	return course, err
}



















































































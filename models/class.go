package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Class struct {
	ID int64				`json:"id"`
	Name string				`json:"name"`
	CourseID int64			`json:"course_id"`
	Plan *string
	Deleted bool
}

func (Class) TableName() string {
	return "scholl.class"
}
func init() {
	dsn := "host=localhost user=postgres password=abc123 dbname=myDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	var class Class
	db.First(&class)
	fmt.Println(class)
}

func GetAllClasses () (classes []Class, err error) {
	//get all -> return ve list
	db, err := setDatabase()
	if err == nil {
		err = db.Where("deleted = ?", false).Find(&classes).Error	//cai nao lq toi db thi                          viet theo ten cot trong db
	}
	return
}

func GetClassByID (IDClass int64) (class Class, err error) {
	//get by id, chi can 1 -> return 1 student thui
	db, err := setDatabase()
	if err == nil {
		err = db.Where("id = ? AND deleted = ?", IDClass, false).Find(&class).Error
	}
	return
}

func AddOneClass (class Class) (Class, error) {   	//output chi can biet variable type. sau do phai cho biet return cai gi
	db, err := setDatabase()
	if err == nil {
		err = db.Create(&class).Error   //dung info nhap vao de creat new student roi tra ve co full info
	}
	return class, err
}

func UpdateOneClass (IDClass int64, newClass Class) (Class, error) {
	db, err := setDatabase()
	if err == nil {
		oldClass, err := GetClassByID(IDClass) //ko can gan "student, err := GetStudentByID(IDStudent)" vi o input da co
		if err == nil && oldClass.ID != 0 {
			if newClass.Name != "" {
				oldClass.Name = newClass.Name
			}
			if newClass.CourseID > 0 {
				oldClass.CourseID = newClass.CourseID
			}
			err = db.Save(&oldClass).Error
			newClass = oldClass
		}
	} else {
		fmt.Println("No data :<")
	}
	return newClass, err
}

func DeleteOneClass (IDClass int64) (Class, error) {
	db, err := setDatabase()
	var class Class
	if err == nil {
		class, err = GetClassByID(IDClass) 	//gan vo chu ai biet dau ma tim
		if err == nil && class.ID != 0 {
			class.Deleted = true			//chi soft delete. ko dc db.delete&
			db.Save(&class)		//update do database
		}
	} else {
		fmt.Println("No data :<")
	}
	return class, err
}
















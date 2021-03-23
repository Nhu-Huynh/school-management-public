package models

import (
	"fmt"
	"time"
)

type StudentClass struct {
	ID int64				`json:"id"`
	StudentID int64		`json:"student_id"`
	ClassID int64			`json:"class_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted bool
}
func (StudentClass) TableName() string {
	return "scholl.student_class"
}
func init() {
	//dsn := "host=localhost user=postgres password=abc123 dbname=myDB port=5432 sslmode=disable"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//fmt.Println(db, err)
	//var studentClass StudentClass
	//db.First(&studentClass)
	//fmt.Println(studentClass)
	//s, e := GetClassByStudent(1)
	//fmt.Println(e)
	//fmt.Println(s)
}



func GetAllStudentClasses () (studentclasses []StudentClass, err error) {
	//get all -> return ve list
	db, err := setDatabase()
	if err == nil {
		err = db.Where("deleted = ?", false).Find(&studentclasses).Error	//cai nao lq toi db thi                          viet theo ten cot trong db

	}

	return
}

func GetStudentClassByID (IDStudentClass int64) (studentclass StudentClass, err error) {
	//get by id, chi can 1 -> return 1 student thui
	db, err := setDatabase()
	if err == nil {
		err = db.Where("id = ? AND deleted = ?", IDStudentClass, false).Find(&studentclass).Error
	}
	return
}

func GetClassByStudent (IDStudent int64) (classes []Class, err error) {		//tim hs hoc lop nao
	//get by id, chi can 1 -> return 1 student thui
	db, err := setDatabase()
	if err == nil {
		err = db.Debug().Select("c.*").Table("scholl.student_class AS sc").Joins("JOIN scholl.class AS c ON sc.class_id = c.id").Where("sc.student_id = ?", IDStudent).Find(&classes).Error
	}
	return
}

func GetStudentByClass (IDClass int64) (students []Student, err error) {
	db, err := setDatabase()
	if err == nil {
		err = db.Debug().Select("s.*").Table("scholl.student_class AS sc").Joins("JOIN scholl.students AS s ON sc.student_id = s.id").Where("sc.class_id = ?", IDClass).Find(&students).Error
	}
	return
}


func AddOneStudentClass (studentclass StudentClass) (StudentClass, error) {   	//output chi can biet variable type. sau do phai cho biet return cai gi
	db, err := setDatabase()
	if err == nil {
		err = db.Create(&studentclass).Error   //dung info nhap vao de creat new student roi tra ve co full info
	}
	return studentclass, err
}

func UpdateOneStudentClass (IDStudentClass int64, newStudentClass StudentClass) (StudentClass, error) {
	db, err := setDatabase()
	if err == nil {
		oldStudentClass, err := GetStudentClassByID(IDStudentClass) //ko can gan "student, err := GetStudentByID(IDStudent)" vi o input da co
		if err == nil && oldStudentClass.ID != 0 {
			oldStudentClass.ID = newStudentClass.ID
			err = db.Save(&oldStudentClass).Error
			newStudentClass = oldStudentClass
		}
	} else {
		fmt.Println("No data :<")
	}
	return newStudentClass, err
}

func DeleteOneStudentClass (IDStudentClass int64) (StudentClass, error) {
	db, err := setDatabase()
	var studentclass StudentClass
	if err == nil {
		studentclass, err = GetStudentClassByID(IDStudentClass) 	//gan vo chu ai biet dau ma tim
		if err == nil && studentclass.ID != 0 {
			studentclass.Deleted = true
			db.Save(&studentclass)		//update vao database
		}
	} else {
		fmt.Println("No data :<")
	}
	return studentclass, err
}























































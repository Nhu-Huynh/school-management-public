package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID int64					`json:"id"`
	Name string					`json:"name"`
	Phone int64					`json:"phone"`
	DOB string					`json:"dob"`
	Deleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Student) TableName() string {
	return "scholl.students"
}
func init() {
	dsn := "host=localhost user=postgres password=abc123 dbname=myDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	//var students []Student
	//db.First(&student)
	//db.Find(&students)
	//db.Last(&students)
	//db.Where("id = ?", 1).Find(&students)

	//update 1 student
	//student1 := Student {Name: "JayKay", Phone: 11992288 , DOB: "1997-09-01"}
	//_ = db.Create(&student1)
	//fmt.Println("result")

	//update nhieu student
	//var student2 = []Student {{Name: "RJay", Phone: 119988 , DOB: "1992-05-12"}, {Name: "Koya", Phone: 992288 , DOB: "1994-09-12"}, {Name: "BonBon", Phone: 112288 , DOB: "1999-03-06"}}
	//db.Create(&student2)

//update student
//	db.Find(&student, []int{8})
//
//	student.Name = "jinzhu 1"
//	student.DOB = "2000-04-09"
//	db.Save(&student)
//

//delete
//	db.Where("Name = ?", "jinzhu 1").Delete(&student)

	//find students have DOB in Aug
	//var students []Student
	//db.Where("EXTRACT(MONTH FROM dob) = ?", 8).Find(&students)
	//fmt.Println(students)

	//change month 08 to 01
	//for i, _ := range students {
	//	students[i].DOB = strings.Replace(students[i].DOB, "-08-", "-01-", 1)
	//}
	//	db.Save(&students)

	//delete phone who have month 01
	//var students []Student
	//db.Where("EXTRACT(MONTH FROM dob) = ?", 1).Find(&students)
	//for i, _ := range students {
	//	students[i].Phone =  0
	//}
	//db.Save(&students)

	//delete id > 3
	//var students []Student
	//db.Where("id > ? AND deleted = ?", 3, false).Find(&students)
	//for i, _ := range students {
	//		students[i].Deleted =  true
	//}
	//
	//db.Save(&students)
	//students, _ := DeleteOneStudent(4)
	//fmt.Println(students)
}

func setDatabase() (db *gorm.DB, err error ) {
	dsn := "host=localhost user=postgres password=abc123 dbname=myDB port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}
func GetAllStudents () (students []Student, err error) {
	//get all -> return ve list
	db, err := setDatabase()
	if err == nil {
		err = db.Where("deleted = ?", false).Find(&students).Error
	}
	return
}

func GetStudentByID (IDStudent int64) (student Student, err error) {
	//get by id, chi can 1 -> return 1 student thui
	db, err := setDatabase()
	if err == nil {
		err = db.Where("id = ? AND deleted = ?", IDStudent, false).Find(&student).Error
	}
	return
}

func AddOneStudent (student Student) (Student, error) {   	//output chi can biet variable type. sau do phai cho biet return cai gi
	db, err := setDatabase()
	if err == nil {
		//students := Student {Name: "NameStudent", Phone: PhoneStudent , DOB: "DOBStudent"}
		_ = db.Create(&student)   //dung info nhap vao de creat new student roi tra ve co full info
	}
	return student, err
}


func UpdateOneStudent (IDStudent int64, newStudent Student) (Student, error) {
	db, err := setDatabase()
	if err == nil {
		oldStudent, err := GetStudentByID(IDStudent) //ko can gan "student, err := GetStudentByID(IDStudent)" vi o input da co
		if err == nil && oldStudent.ID != 0 {
			oldStudent.Name = newStudent.Name
			err = db.Save(&oldStudent).Error
			newStudent = oldStudent
		}
	} else {
		fmt.Println("No data :<")
	}
	return newStudent, err
}

func DeleteOneStudent (IDStudent int64) (Student, error) {
	db, err := setDatabase()
	var student Student
	if err == nil {
		student, err = GetStudentByID(IDStudent) 	//gan vo chu ai biet dau ma tim
		if err == nil && student.ID != 0 {
			student.Deleted = true
			db.Save(&student)		//update vao database
		}
	} else {
		fmt.Println("No data :<")
	}
	return student, err
}

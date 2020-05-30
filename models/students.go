package models

import "github.com/jinzhu/gorm"

//Student ...
type Student struct {
	gorm.Model
	Name           string `json:"name"`
	Age            string `json:"age"`
	Standard       string `json:"standard"`
	SchoolCode     int    `json:"schoolCode"`
	IdentityNumber int    `json:"identityNumber"`
}

//StudentService ...
type StudentService interface {
	AddStudent(input Student) (*Student, error)
	GetStudent(id int) (*Student, error)
	GetStudents(schoolCode int) ([]Student, error)
	DeleteStudent(id int) (*Student, error)
	UpdateStudent(id int, input Student) (*Student, error)
}

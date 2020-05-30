package models

import "github.com/tyagip966/common-repo/models"

//StudentService ...
type StudentService interface {
	AddStudent(input models.Student) (*models.Student, error)
	GetStudent(id int) (*models.Student, error)
	GetStudents(schoolCode int) ([]models.Student, error)
	DeleteStudent(id int) (*models.Student, error)
	UpdateStudent(id int, input models.Student) (*models.Student, error)
}

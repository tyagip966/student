package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/tyagip966/common-repo/models"
)

//StudentService ...
type StudentService struct {
	Database *gorm.DB
}

//AddStudent ...
func (s StudentService) AddStudent(input models.Student) (*models.Student, error) {
	db := s.Database.Create(input)
	if db.Error != nil {
		return nil, db.Error
	}
	return &input, nil
}

//GetStudent ...
func (s StudentService) GetStudent(id int) (*models.Student, error) {
	var response models.Student
	db := s.Database.Model(&models.Student{}).Where("id = ?", id).Find(&response)
	if db.Error != nil {
		return nil, db.Error
	}
	return &response, nil
}

//GetStudents ...
func (s StudentService) GetStudents(schoolCode int) ([]models.Student, error) {
	var response []models.Student
	db := s.Database.Model(&models.Student{}).Where("school_code = ?", schoolCode).Find(&response)
	if db.Error != nil {
		return nil, db.Error
	}
	return response, nil
}

//DeleteStudent ...
func (s StudentService) DeleteStudent(id int) (*models.Student, error) {
	var response models.Student
	db := s.Database.Model(&models.Student{}).Where("id = ?", id).Delete(&models.Student{})
	if db.Error != nil {
		return nil, db.Error
	}
	return &response, nil
}

//UpdateStudent ...
func (s StudentService) UpdateStudent(id int, input models.Student) (*models.Student, error) {
	return nil, nil
}

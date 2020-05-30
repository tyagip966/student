package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"student/models"
)

func startUp(dns string) (*gorm.DB, error) {
	return connection(dns)
}

func connection(dns string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dns)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to Database !!")
	db.SingularTable(true)
	db.DropTableIfExists("students")
	db.AutoMigrate(&models.Student{})
	return db, nil
}

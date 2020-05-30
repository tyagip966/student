package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"student"
	"student/constants"
	"student/models/postgres"
)

type Container struct{
    Server    *student.GrpcServer
    Service   *postgres.StudentService
    DB        *gorm.DB
	Injected                 bool
}
func (c *Container) TriggerDI() error{
	if c.Injected {
		return nil
	}
	err := c.StartServer()
	if err != nil {
		c.Injected = false
		return err
	}
	c.Injected = true
	return nil
}

func (c *Container) StartServer() error{
  if c.Server == nil {
	  var err error
	  c.Server.Service,err = student.ListenGRPC(*c.GetService(),constants.ServerPort)
  	 if err != nil {
  	 	return err
	 }
  }
  return nil
}

func (c *Container) GetService() *postgres.StudentService {
     if c.Service == nil {
     	c.Service = &postgres.StudentService{
     		Database: c.GetDB(),
		}
	 }
	 return c.Service
}

func (c *Container) GetDB() *gorm.DB{
    if c.DB == nil {
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", constants.UERNAME, constants.PASSWORD, constants.HOST, constants.PORT, constants.DBNAME)
    	db,err  := startUp(dsn)
    	if err != nil {
    		log.Fatal("Error Connecting to DB",err)
		}
		c.DB = db
	}

	return c.DB
}
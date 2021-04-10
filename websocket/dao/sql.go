package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
var DB *gorm.DB
func MysqlInit(){
	dsn := "root:dengjie123@tcp(127.0.0.1:3306)/web?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Panicln(err)
	}
}
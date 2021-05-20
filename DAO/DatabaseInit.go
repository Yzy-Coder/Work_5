package DAO

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"Work_5/object"
)


func DataBaseInit() (*gorm.DB, object.ErrMessage) {

	DB , err := gorm.Open("mysql","root:@(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local")

	return DB , object.ErrMessage{Error: err}
}
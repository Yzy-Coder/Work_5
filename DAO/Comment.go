package DAO

import (
	"Work_5/object"
	"github.com/jinzhu/gorm"
)

//创建评论记录
func CreateComment(comment *object.Comment, db *gorm.DB) object.ErrMessage {

	//检查评论信息表是否存在，不存在则创建
	if !db.HasTable(&comment) {
		db.AutoMigrate(&comment)
	}

	//创建相应评论记录
	db.Create(&comment)
	return object.ErrMessage{}
}

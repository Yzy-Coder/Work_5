package DAO

import (
	"Work_5/object"
	"github.com/jinzhu/gorm"
)

//创建点赞映射关系
func CreateArticleThumbsUpMapping(mapping *object.ArticleThumbsUpMapping, db *gorm.DB) object.ErrMessage {

	//检查映射表是否存在，不存在则创建
	if !db.HasTable(&mapping) {
		db.AutoMigrate(&mapping)
	}

	//插入信息
	db.Create(&mapping)
	return object.ErrMessage{}
}

//删除点赞映射关系
func DeletArticleThumbsUpMapping(mapping *object.ArticleThumbsUpMapping, db *gorm.DB) object.ErrMessage {

	result := db.Delete(&mapping)
	if result.RowsAffected > 0 {
		return object.ErrMessage{}
	} else {
		return object.ErrMessage{IsErr: true, Whaterror: "Delet err"}
	}
}

//查询点赞映射关系是否存在
func QueryArticleThumbsUpMapping(mapping *object.ArticleThumbsUpMapping, db *gorm.DB) (bool, object.ErrMessage) {

	result := db.Where("user_id = ? AND article_id = ?", mapping.UserId, mapping.ArticleId).Find(&mapping)
	if result.RowsAffected > 0 {
		return true, object.ErrMessage{}
	} else {
		return false, object.ErrMessage{}
	}
}

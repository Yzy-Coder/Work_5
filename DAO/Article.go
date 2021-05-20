package DAO

import (
	"Work_5/object"
	"github.com/jinzhu/gorm"
)

//创建新的文章
func CreateArticle(article *object.Article, db *gorm.DB) object.ErrMessage {

	if !db.HasTable(&article) {
		db.AutoMigrate(&article)
	}

	db.Create(&article)
	return object.ErrMessage{}
}

//查询文章是否存在
func ArticleQuerry(article *object.Article, db *gorm.DB) (bool, object.ErrMessage) {

	result := db.Where("id = ?", article.ID).Find(&article)

	if result.RowsAffected > 0 {
		return true, object.ErrMessage{}
	} else {
		return false, object.ErrMessage{IsErr: true, Whaterror: "article not exist"}
	}
}

//文章点赞数+1
func ThumbsNumUp(article *object.Article, db *gorm.DB) object.ErrMessage {

	db.Model(&article).Update("thumbs_up_num", gorm.Expr("thumbs_up_num + ?", 1))
	return object.ErrMessage{}
}

//文章点赞数-1
func ThumbsNumDown(article *object.Article, db *gorm.DB) object.ErrMessage {

	db.Model(&article).Update("thumbs_up_num", gorm.Expr("thumbs_up_num - ?", 1))
	return object.ErrMessage{}
}

//将文章按点赞排名前十排序并返回
func SelectTopArticle(articles *[]object.Article, db *gorm.DB) object.ErrMessage {

	//将查询结果存放到article切片
	db.Model(&object.Article{}).Order("thumbs_up_num desc").Find(&articles).Limit(10)

	return object.ErrMessage{}
}

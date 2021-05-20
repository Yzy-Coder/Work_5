package DAO

import (
	"Work_5/object"
	"github.com/jinzhu/gorm"
)

//提问
func CreateQuestion(question *object.Question, db *gorm.DB) object.ErrMessage {

	//判断是否存在相应的表，不存在则创建
	if !db.HasTable(&object.Question{}) {
		db.AutoMigrate(&object.Question{})
	}

	//插入问题信息
	db.Create(&question)
	return object.ErrMessage{}
}

//查询问题是否存在
func QuerryQuestion(question *object.Question, db *gorm.DB) (bool, object.ErrMessage) {

	//查询问题是否存在
	result := db.Where("id = ?", question.ID).Find(&question)

	//存在返回true,不存在返回false并报错
	if result.RowsAffected > 0 {
		return true, object.ErrMessage{}
	} else {
		return false, object.ErrMessage{IsErr: true, Whaterror: "Question not exist"}
	}

}

//判断用户是否为提问人
func IsQuestioner(user *object.User, question object.Question, db *gorm.DB) (bool, object.ErrMessage) {

	//根据id查找问题信息
	db.Find(&question)

	//判断是否是提问人
	if question.QuestionerId != user.UserId {
		return false, object.ErrMessage{}
	} else {
		return true, object.ErrMessage{}
	}

}

//更新问题信息
func UpdateQuestion(question *object.Question, db *gorm.DB) object.ErrMessage {

	//更新语句
	result := db.Model(&question).Update("question_content", question.QuestionContent)

	//判断更新是否成功，失败返回错误信息
	if result.RowsAffected > 0 {
		return object.ErrMessage{}
	} else {
		return object.ErrMessage{IsErr: true, Whaterror: "update false"}
	}
}

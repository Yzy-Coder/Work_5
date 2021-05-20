package Service

import (
	"Work_5/DAO"
	"Work_5/object"
)

//回答问题（发表文章）方法
func AnswerQuestion(user *object.User, question *object.Question, article *object.Article) object.ErrMessage {

	//获取数据库连接
	db, err := DAO.DataBaseInit()
	if err.IsErr == true {
		return err
	}

	//判断用户和问题是否存在
	userexist, _ := DAO.UserQuery(user, db)
	questionexist, _ := DAO.QuerryQuestion(question, db)

	//存在则创建文章信息，回答问题
	if userexist && questionexist {
		DAO.CreateArticle(article, db)
		return object.ErrMessage{}
	}

	//不存在用户或者提问，返回错误信息
	return object.ErrMessage{IsErr: true, Whaterror: "user or question not exist"}
}

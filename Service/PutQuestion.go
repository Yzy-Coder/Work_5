package Service

import (
	"Work_5/DAO"
	"Work_5/object"
)

//提问方法
func PutQuestion(user *object.User, question *object.Question) object.ErrMessage {

	//获取数据库连接
	db, err := DAO.DataBaseInit()
	if err.IsErr {
		return err
	}

	//查询用户是否存在
	isexist, _ := DAO.UserQuery(user, db)
	if !isexist {
		err.IsErr = true
		err.Whaterror = "user not exist"
		return err
	}

	//如果用户存在，创建相应记录

	//这一行貌似可以删了，因为question的QuestionerId也接受userid这个表单信息
	question.QuestionerId = user.UserId

	//调用DAO提出问题的方法
	err = DAO.CreateQuestion(question, db)
	if err.IsErr {
		return err
	} else {
		return object.ErrMessage{}
	}

}

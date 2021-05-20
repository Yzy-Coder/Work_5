package Service

import (
	"Work_5/DAO"
	"Work_5/object"
)

//Service修改提问函数
func UpdataQuestion(user *object.User, question *object.Question) object.ErrMessage {

	//获得数据库连接
	db, err := DAO.DataBaseInit()
	if err.IsErr {
		return err
	}

	//判断修改提问的用户，是否是这个问题的提出者
	isQuestioner, _ := DAO.IsQuestioner(user, *question, db)
	if !isQuestioner {
		return object.ErrMessage{IsErr: true, Whaterror: "user not questioner"}
	}

	//调用DAO修改提问的方法
	err = DAO.UpdateQuestion(question, db)

	if err.IsErr {
		return err
	} else {
		return object.ErrMessage{}
	}

}

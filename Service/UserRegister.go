package Service

import (
	"Work_5/DAO"
	"Work_5/object"
)

//用户注册
func UserRegister(user *object.User) object.ErrMessage {

	//获得数据库连接
	DB, err := DAO.DataBaseInit()
	if err.IsErr == true {
		return err
	}

	//检查用户名是否重复，如果重复返回错误信息
	var isexist bool
	isexist, err = DAO.UserNameQuery(user, DB)
	if isexist == true {
		return err
	}

	//用户名不重复，调用DAO方法
	err = DAO.UserRegister(user, DB)
	return err
}

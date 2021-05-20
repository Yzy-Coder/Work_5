package Service

import (
	"Work_5/DAO"
	"Work_5/object"
)

//用户登录
func UserLogin(user *object.User) object.ErrMessage {

	//获得数据库连接
	db, err := DAO.DataBaseInit()
	if err.IsErr {
		return err
	}

	//创建existuser用于判断用户是否存在以及后续密码校验
	var existuser object.User
	existuser.UserId = user.UserId

	//查询用户是否存在，不存在则返回相应的错误信息
	isexist, err := DAO.UserQuery(&existuser, db)
	if !isexist {
		return object.ErrMessage{IsErr: true, Whaterror: "user not exist"}
	}

	//校验用户密码是否正确，不成却则报错
	if user.UserPassword == existuser.UserPassword {

		//将用户名传给user
		user.UserName = existuser.UserName
		return object.ErrMessage{}
	} else {
		return object.ErrMessage{IsErr: true, Whaterror: "PasswordWrong"}
	}

}

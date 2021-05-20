package DAO

import (
	"Work_5/object"
	"github.com/jinzhu/gorm"
)

//用户注册方法
func UserRegister(user *object.User, db *gorm.DB) object.ErrMessage {

	//检查对应的用户信息表是否存在，不存在则创建
	if !db.HasTable(&object.User{}) {
		db.AutoMigrate(&object.User{})
	}

	//插入用户信息
	result := db.Create(&user)

	//检查插入是否成功，失败则返回相关错误信息
	if result.RowsAffected > 0 {
		return object.ErrMessage{IsErr: false}
	} else {
		return object.ErrMessage{IsErr: true, Whaterror: "Register faild"}
	}

}

//查询用户是否存在（根据ID）
func UserQuery(user *object.User, db *gorm.DB) (bool, object.ErrMessage) {

	//查询用户是否存在,存在返回true,不存在返回false
	result := db.Where("user_id = ?", user.UserId).Find(&user)
	if result.RowsAffected > 0 {
		return true, object.ErrMessage{}
	} else {
		return false, object.ErrMessage{}
	}

}

//查询用户名是否存在
func UserNameQuery(user *object.User, db *gorm.DB) (bool, object.ErrMessage) {

	//查询用户名是否存在，存在则返回true，并返回错误信息（因为不支持用户名重复），不存在返回false
	result := db.Where("user_name = ?", user.UserName).Find(&user)
	if result.RowsAffected > 0 {
		return true, object.ErrMessage{IsErr: true, Whaterror: "username repeat"}
	} else {
		return false, object.ErrMessage{}
	}

}

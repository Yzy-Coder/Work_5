package API

import (
	"Work_5/Service"
	"Work_5/object"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户注册API方法（检测用户名是否重复，不支持重复）
func UserRegister(ctx *gin.Context) {

	//绑定用户数据
	var user object.User
	ctx.ShouldBind(&user)

	//调用用户注册Service方法
	err := Service.UserRegister(&user)

	//注册成功返回用户结构体，失败返回错误信息
	if err.IsErr == true {
		ctx.JSON(http.StatusOK, err)

	} else {
		ctx.JSON(http.StatusOK, user)
	}

}

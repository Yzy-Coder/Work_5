package API

import (
	"Work_5/Service"
	"Work_5/object"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户登录API方法
func UserLogin(ctx *gin.Context) {

	//绑定user结构体
	var user object.User
	ctx.ShouldBind(&user)

	//调用Service层用户登录方法，并接受err
	err := Service.UserLogin(&user)

	//若成功则返回用户信息，否则返回err信息
	if err.IsErr {
		ctx.JSON(http.StatusOK, err)
	} else {
		ctx.JSON(http.StatusOK, user)
	}

}

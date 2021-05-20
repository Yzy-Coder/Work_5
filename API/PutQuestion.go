package API

import (
	"Work_5/Service"
	"Work_5/object"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutQuestion(ctx *gin.Context) {

	//绑定user和question结构体
	var user object.User
	var question object.Question
	ctx.ShouldBind(&user)
	ctx.ShouldBind(&question)

	//调用Service层提问的方法
	err := Service.PutQuestion(&user, &question)

	//根据调用结果返回
	if err.IsErr {
		ctx.JSON(http.StatusOK, err)
	} else {
		ctx.JSON(http.StatusOK, question)
	}

}

package API

import (
	"Work_5/Service"
	"Work_5/object"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateQuestion(ctx *gin.Context) {

	//绑定user，question结构体
	var user object.User
	var question object.Question
	ctx.ShouldBind(&user)
	ctx.ShouldBind(&question)

	//调用Service的方法
	err := Service.UpdataQuestion(&user, &question)

	//根据调用结果返回相关信息，失败则返回错误信息，成功返回修改后的提问信息
	if err.IsErr {
		ctx.JSON(http.StatusOK, err)
	} else {
		ctx.JSON(http.StatusOK, question)
	}
}

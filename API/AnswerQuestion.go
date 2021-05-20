package API

import (
	"Work_5/Service"
	"Work_5/object"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AnswerQuestion(ctx *gin.Context) {

	//绑定user，question，article参数
	var user object.User
	var question object.Question
	var article object.Article
	ctx.ShouldBind(&user)
	ctx.ShouldBind(&question)
	ctx.ShouldBind(&article)

	//调用发表文章（回复问题）的Service方法
	error := Service.AnswerQuestion(&user, &question, &article)

	//根据调用结果返回
	if error.IsErr {
		ctx.JSON(http.StatusOK, error)
	} else {
		ctx.JSON(http.StatusOK, article)
	}

}

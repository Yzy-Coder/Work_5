package API

import (
	"Work_5/Service"
	"Work_5/object"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Comment(ctx *gin.Context) {

	//绑定user，article，comment的参数
	var user object.User
	var article object.Article
	var comment object.Comment
	ctx.ShouldBind(&user)
	ctx.ShouldBind(&article)
	ctx.ShouldBind(&comment)

	//调用评论文章的Service方法
	err := Service.CommentArticle(&user, &article, &comment)

	//根据调用结果返回
	if err.IsErr {
		ctx.JSON(http.StatusOK, err)
	} else {
		ctx.JSON(http.StatusOK, comment)
	}
}

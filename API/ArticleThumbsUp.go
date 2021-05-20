package API

import (
	"Work_5/Service"
	"Work_5/object"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleThumbs(ctx *gin.Context) {

	//绑定user，article参数
	var user object.User
	var article object.Article
	ctx.ShouldBind(&user)
	ctx.ShouldBind(&article)

	//调用文章点赞的Service方法
	err := Service.ArticleThumbsUp(&user, &article)

	//根据调用结果返回
	if err.IsErr {
		ctx.JSON(http.StatusOK, err)
	} else {
		ctx.JSON(http.StatusOK, "操作成功")
	}
}

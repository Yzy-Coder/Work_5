package API

import (
	"Work_5/Service"
	"Work_5/object"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SelectTopArticle(ctx *gin.Context) {

	//创建article的切片
	var articles []object.Article

	//调用Service层的方法
	err := Service.SelectTopArticle(&articles)

	//根据调用结果返回文章信息或错误信息
	if err.IsErr {
		ctx.JSON(http.StatusOK, err)
	} else {
		ctx.JSON(http.StatusOK, articles)
	}

}

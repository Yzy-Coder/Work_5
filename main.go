package main

import (
	"Work_5/API"
	"github.com/gin-gonic/gin"
)

func main() {

	Route := gin.Default()

	//用户注册
	Route.POST("/userregister", API.UserRegister)

	//用户登录
	Route.POST("/UserLogin", API.UserLogin)

	//提问
	Route.POST("/PutQuestion", API.PutQuestion)

	//修改问题
	Route.POST("/updatequestion", API.UpdateQuestion)

	//回答问题，发布文章
	Route.POST("/answerquestion", API.AnswerQuestion)

	//评论文章
	Route.POST("/comment", API.Comment)

	//点赞文章
	Route.POST("/articlethumbsup", API.ArticleThumbs)

	//按点赞数进行文章排序显示
	Route.POST("/selecttoparticle", API.SelectTopArticle)

	Route.Run(":9090")
}

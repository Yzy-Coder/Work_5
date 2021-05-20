package Service

import (
	"Work_5/DAO"
	"Work_5/object"
)

//评论文章
func CommentArticle(user *object.User, article *object.Article, comment *object.Comment) object.ErrMessage {

	//获取数据库连接
	db, err := DAO.DataBaseInit()
	if err.IsErr {
		return err
	}

	//检查用户和文章是否存在
	userexist, err := DAO.UserQuery(user, db)
	articleexist, err := DAO.ArticleQuerry(article, db)

	//如果存在，则创建评论信息
	if userexist && articleexist {
		DAO.CreateComment(comment, db)
		return object.ErrMessage{}
	}

	//不存在用户或文章则返回错误信息
	return object.ErrMessage{IsErr: true, Whaterror: "user or article not exist"}

}

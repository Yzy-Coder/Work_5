package Service

import (
	"Work_5/DAO"
	"Work_5/object"
)

//文章点赞方法
func ArticleThumbsUp(user *object.User, article *object.Article) object.ErrMessage {

	//获取数据库连接
	db, err := DAO.DataBaseInit()
	if err.IsErr {
		return err
	}

	//检查用户和文章是否存在
	userexist, _ := DAO.UserQuery(user, db)
	articleexist, _ := DAO.ArticleQuerry(article, db)

	//若存在用户和文章
	if userexist && articleexist {

		//创建用户与文章的点赞映射关系
		var mapping = object.ArticleThumbsUpMapping{UserId: user.UserId, ArticleId: article.ID}

		//查询该映射关系是否存在
		mappingexist, _ := DAO.QueryArticleThumbsUpMapping(&mapping, db)

		//映射关系存在则删除映射关系并文章点赞数-1，否则创建映射关系并点赞数+1
		if mappingexist {
			DAO.ThumbsNumDown(article, db)
			DAO.DeletArticleThumbsUpMapping(&mapping, db)
			return object.ErrMessage{}
		} else {
			DAO.ThumbsNumUp(article, db)
			DAO.CreateArticleThumbsUpMapping(&mapping, db)
			return object.ErrMessage{}
		}

	}

	//若不存在文章或用户，则返回错误信息
	return object.ErrMessage{IsErr: true, Whaterror: "user or article not exist"}
}

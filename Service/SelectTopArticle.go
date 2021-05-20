package Service

import (
	"Work_5/DAO"
	"Work_5/object"
)

//获取点赞前十的文章
func SelectTopArticle(articles *[]object.Article) object.ErrMessage {

	//获得数据库连接
	db, err := DAO.DataBaseInit()
	if err.IsErr {
		return err
	}

	//调用DAO层函数
	if err = DAO.SelectTopArticle(articles, db); err.IsErr {
		return err
	}

	return object.ErrMessage{}
}

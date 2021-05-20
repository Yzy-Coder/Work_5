package object

import (
	"time"
)

type ArticleThumbsUpMapping struct {
	Id        int `gorm:"primary_key"`
	CreatedAt time.Time
	UserId    int `form:"userid"`
	ArticleId int `form:"articleid"`
}

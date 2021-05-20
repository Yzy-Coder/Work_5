package object

import (
	"time"
)

type Article struct {
	ID             int `gorm:"primary_key" form:"articleid"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
	AuthorId       int        `form:"userid"`
	QuestionId     int        `form:"questionid"`
	ArticleTitle   string     `form:"articletitle"`
	ArticleContent string     `form:"articlecontent"`
	ThumbsUpNum    int
}

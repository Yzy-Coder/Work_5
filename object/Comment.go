package object

import "time"

type Comment struct {
	ID             uint `gorm:"primary_key" form:"commentid"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
	AuthorId       int        `form:"userid"`
	ArticleId      int        `form:"articleid"`
	CommentContent string     `form:"commentcontent"`
	ThumbsUpNum    int
}

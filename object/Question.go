package object

import "time"

type Question struct {

	ID              int   `form:"questionid" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	QuestionerId    int    `form:"userid"`
	QuestionTitle   string `form:"questiontitle"`
	QuestionContent string `form:"questioncontent"`
	AttentionRate   int
}

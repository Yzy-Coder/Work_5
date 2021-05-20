package object

import "time"

type User struct {
	UserId       int `gorm:"primary_key;AUTO_INCREMENT" form:"userid"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
	UserName     string     `form:"username"`
	UserPassword string     `form:"userpassword"`
}

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Followers []*User   `json:"followers" gorm:"many2many:user_followers"`
	Following []*User   `json:"following" gorm:"many2many:user_following"`
	Posts     []Post    `json:"posts"`
	Comments  []Comment `json:"comments"`
	Profile   Profile   `json:"profile"`
}

type Profile struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	Bio    string `json:"bio"`
	UserID uint
}

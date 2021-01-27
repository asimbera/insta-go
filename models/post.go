package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Photos   string    `json:"photos"`
	Caption  string    `json:"caption"`
	Comments []Comment `json:"comments"`
	UserID   uint
}

type Comment struct {
	gorm.Model
	Text   string `json:"text"`
	PostID uint
	UserID uint
}

package entities

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `gorm:"size:100;not null" json:"title"`
	Content  string `gorm:"type:text;not null" json:"content"`
	AuthorID uint   `gorm:"not null" json:"authorid"`
	Author   User   `gorm:"foreignKey:AuthorID"`
}

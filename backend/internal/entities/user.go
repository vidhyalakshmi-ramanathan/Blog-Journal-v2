package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"size:100;not null" json:"username"`
	Email       string `gorm:"size:100;unique;not null" json:"email"`
	Password    string `gorm:"size:100;not null" json:"-"`
	PhoneNumber string `gorm:"size:50;not null" json:"phonenumber"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint   `gorm:"primaryKey";autoIncrement:true;unique;`
	Username  string `json:"username";binding:"required"`
	Password  string `json:"password";binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Post      []Post         `gorm:"foreignKey:UserID;references:Id;constraint:OnDelete:CASCADE;"`
}

type Post struct {
	gorm.Model
	Id          uint   `gorm:"primaryKey";autoIncrement:true;unique;`
	Title       string `json:"title";binding:"required"`
	Description string `json:"username";binding:"required"`
	Detail      string `json:"password";binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

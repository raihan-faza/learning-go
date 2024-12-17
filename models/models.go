package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint `gorm:"primaryKey";autoIncrement:true;unique;`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Post      []Post         `gorm:"foreignKey:UserID;references:Id;constraint:OnDelete:CASCADE;"`
}

type Post struct {
	gorm.Model
	Id          uint `gorm:"primaryKey";autoIncrement:true;unique;`
	Title       string
	Description string
	Detail      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

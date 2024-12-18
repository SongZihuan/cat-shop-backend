package model

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Show bool   `gorm:"type:boolean;not null"`
}

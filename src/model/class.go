package model

import (
	"github.com/SuperH-0630/cat-shop-back/src/model/modeltype"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Show bool   `gorm:"type:boolean;not null"` // 仅不展示
	Down bool   `gorm:"type:boolean;not null"` // 下架所有商品
}

func (*Class) TableName() string {
	return "class"
}

func NewEmptyClass() *Class {
	return &Class{
		Model: gorm.Model{
			ID: modeltype.ClassEmptyID,
		},
		Name: modeltype.ClassEmptyName,
		Show: modeltype.ClassEmptyShow,
	}
}

func (cls *Class) ResetEmpty() {
	if !cls.IsEmpty() {
		panic("class is not empty")
	}
	cls.Name = modeltype.ClassEmptyName
	cls.Show = modeltype.ClassEmptyShow
}

func (cls *Class) IsEmpty() bool {
	return cls.ID == modeltype.ClassEmptyID
}

func (cls *Class) IsEmptyWithCheck() bool {
	if !cls.IsEmpty() {
		return false
	}

	return cls.Show == modeltype.ClassEmptyShow && cls.Name == modeltype.ClassEmptyName
}

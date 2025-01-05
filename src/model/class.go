package model

import (
	"github.com/SongZihuan/cat-shop-backend/src/model/modeltype"
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
		Down: modeltype.ClassEmptyDown,
	}
}

func (cls *Class) resetEmpty() bool {
	if !cls.IsEmpty() {
		panic("class is not empty")
	}

	if cls.Name == modeltype.ClassEmptyName && cls.Show == modeltype.ClassEmptyShow && cls.Down == modeltype.ClassEmptyDown {
		return false
	}

	cls.Name = modeltype.ClassEmptyName
	cls.Show = modeltype.ClassEmptyShow
	cls.Down = modeltype.ClassEmptyDown
	return true
}

func (cls *Class) IsEmpty() bool {
	return cls.ID == modeltype.ClassEmptyID
}

func (cls *Class) ResetIsEmpty() bool {
	if !cls.IsEmpty() {
		return false
	}

	return cls.resetEmpty()
}

func (cls *Class) IsClassDown() bool {

	return !cls.IsEmpty() && cls.Down
}

func (cls *Class) IsClassDownOrNotShow() bool {
	return cls.IsEmpty() || !cls.Show || cls.IsClassDown()
}

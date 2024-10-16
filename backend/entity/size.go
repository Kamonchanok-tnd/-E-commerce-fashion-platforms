package entity

import "gorm.io/gorm"

type Size struct {
	gorm.Model
	SizeName	string
}
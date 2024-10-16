package entity

import "gorm.io/gorm"

type Color struct {
	gorm.Model
	ColorName	string
}
package entity

import "gorm.io/gorm"

type Target struct {
	gorm.Model
	TargetName	string

	Product []Product  `gorm:"foreignKey:TargetID"`
}
package entity

import "gorm.io/gorm"

type Delivery struct {
	gorm.Model
	Tracking	string

	UserID   uint
	CartID   uint
	StatusID   uint

	User   User    `gorm:"foreignKey:UserID"`
	Cart   Cart    `gorm:"foreignKey:CartID"`
	Status   Status    `gorm:"foreignKey:StatusID"`
}
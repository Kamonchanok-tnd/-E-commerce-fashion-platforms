package entity

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	TotalPrice	float64

	CartID   uint
	StatusID   uint

	Cart   Cart    `gorm:"foreignKey:CartID"`
	Status   Status    `gorm:"foreignKey:StatusID"`
}
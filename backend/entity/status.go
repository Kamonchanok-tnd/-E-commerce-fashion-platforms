package entity

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	StatusName	string
	Type		string

	Cart []Cart `gorm:"foreignKey:StatusID"`
	Product []Product `gorm:"foreignKey:StatusID"`
	Delivery []Delivery `gorm:"foreignKey:StatusID"`
	Payment []Payment `gorm:"foreignKey:StatusID"`
}
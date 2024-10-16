package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Amount	int

	UserID   uint
	ProductID   uint
	StatusID   uint

	User   User    `gorm:"foreignKey:UserID"`
	Product   Product    `gorm:"foreignKey:ProductID"`
	Status   Status    `gorm:"foreignKey:StatusID"`
}
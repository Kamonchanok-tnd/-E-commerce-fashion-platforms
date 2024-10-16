package entity

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Score 		float64
	Comment		string
	
	ProductID   uint
	UserID   uint

	Product 	Product `gorm:"foreignKey:ProductID"`
	User		User	 `gorm:"foreignKey:UserID"`
}
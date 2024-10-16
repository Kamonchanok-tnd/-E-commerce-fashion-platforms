package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName   string
	Email      string
	Password   string
	Address string
	Phone   string
	Role       string

	Cart []Cart `gorm:"foreignKey:UserID"`
	Rating []Rating  `gorm:"foreignKey:UserID"`
	Delivery []Delivery  `gorm:"foreignKey:UserID"`

}
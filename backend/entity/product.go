package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName	string
	Price 		float64
	Amount		int
	Detail		string
	Image 		[]byte  `gorm:"type:blob"`
	
	CategoryID   uint
	SizeID   uint
	ColorID   uint
	StatusID   uint
	TargetID   	uint

	Category 	Category `gorm:"foreignKey:CategoryID"`
	Color		Color	 `gorm:"foreignKey:ColorID"`
	Size		Size	 `gorm:"foreignKey:SizeID"`
	Status   	Status   `gorm:"foreignKey:StatusID"`
	Target		Target	 `gorm:"foreignKey:TargetID"`

	Rating []Rating  `gorm:"foreignKey:ProductID"`
}
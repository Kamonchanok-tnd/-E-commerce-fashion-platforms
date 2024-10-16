package config

import (
	"fmt"

	"github.com/Kamonchanok-tnd/E-commerce-fashion-platforms/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("shop.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {

	// AutoMigrate สำหรับทุก entity
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Category{},
		&entity.Color{},
		&entity.Size{},
		&entity.Status{}, 
		&entity.Product{},  
		&entity.Cart{},
		&entity.Payment{},
		&entity.Delivery{},
		&entity.Rating{},
	)
	if err != nil {
		fmt.Println("Error in AutoMigrate:", err)
	} else {
		fmt.Println("AutoMigrate completed successfully.")
	}

	//create user
	hashedPassword, _ := HashPassword("123")
	User1 := &entity.User{
		UserName:   "admin",
		Email:      "admin@gmail.com",
		Password:   hashedPassword,
		Role:       "admin",
	}
	db.FirstOrCreate(User1, &entity.User{
		Email: "admin@gmail.com",
	})

	hashedPassword2, _ := HashPassword("123")
	User2 := &entity.User{
		UserName:   "user1",
		Email:      "user@gmail.com",
		Password:   hashedPassword2,
		Phone: 		"0123456789",
		Role:       "user",
	}
	db.FirstOrCreate(User2, &entity.User{
		Email: "user@gmail.com",
	})
}
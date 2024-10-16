package main

import (
	"github.com/Kamonchanok-tnd/E-commerce-fashion-platforms/config"
)

const PORT = "8000"

func main() {

	// open connection database
	config.ConnectionDB()

	// Generate databases
	config.SetupDatabase()

	// ไม่มี router ส่วนนี้คือสำหรับการทดสอบสร้างฐานข้อมูลเท่านั้น
	println("Database setup complete. No routes are running.")
}

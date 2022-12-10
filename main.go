package main

import (
	"log"
	"os"
	"project_api/payment"
	"project_api/product"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(localhost:3306)/bank?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL))
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(product.Product{}, payment.Payment{})
}

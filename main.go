package main

import (
	"log"
	"os"
	"project_api/handler"
	"project_api/payment"
	"project_api/product"

	"github.com/gin-gonic/gin"
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
		log.Println("Connection Failed to Open")
	}
	log.Println("Connection Established")

	db.AutoMigrate(product.Product{}, payment.Payment{})

	pr := product.NewRepository(db)
	ps := product.NewService(pr)
	productHandler := handler.NewProductHandler(ps)

	paymentR := payment.NewRepository(db)
	paymentS := payment.NewService(paymentR)
	paymentHandler := handler.NewPaymentHandler(paymentS)

	router := gin.Default()
	router.GET("/", productHandler.Hello)

	api := router.Group("/api")
	api.POST("/product", productHandler.Store)
	api.PUT("/product/:id/update", productHandler.Update)
	api.DELETE("/product/:id/delete", productHandler.Delete)
	api.GET("/product/:id", productHandler.FetchById)
	api.GET("/products", productHandler.List)

	api.POST("/payment", paymentHandler.Store)
	api.PUT("/payment/:id/update", paymentHandler.Update)
	api.DELETE("/payment/:id/delete", paymentHandler.Delete)
	api.GET("/payment/:id", paymentHandler.FetchById)
	api.GET("/payments", paymentHandler.List)

	router.Run(":3000")

}

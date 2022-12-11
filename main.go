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
		log.Fatalln(err)
	}

	db.AutoMigrate(product.Product{}, payment.Payment{})

	pr := product.NewRepository(db)
	ps := product.NewService(pr)
	productHandler := handler.NewProductHandler(ps)

	router := gin.Default()
	router.GET("/", productHandler.Hello)

	api := router.Group("/api")
	api.POST("/product", productHandler.Store)
	api.PUT("/product/:id/update", productHandler.Update)
	api.DELETE("/product/:id/delete", productHandler.Delete)
	api.GET("/product/:id", productHandler.FetchById)
	api.GET("/products", productHandler.List)

	router.Run(":3000")

}

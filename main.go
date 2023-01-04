package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	user "project_api/User"
	"project_api/adapter"
	docs "project_api/docs"
	"project_api/handler"
	"project_api/middlewares"
	"project_api/payment"
	"project_api/product"
	service "project_api/services"
)

// @title           Project Go API [ESGI]
// @version         1.0
// @description     This is a sample document for an api implementation.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				API works with JWT Bearer Token

func main() {
	godotenv.Load(".env")

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(localhost:3306)/bank?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL))
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	log.Println("Connection Established")

	db.AutoMigrate(product.Product{}, payment.Payment{}, user.User{})

	pr := product.NewRepository(db)
	ps := product.NewService(pr)
	productHandler := handler.NewProductHandler(ps)

	paymentR := payment.NewRepository(db)
	paymentS := payment.NewService(paymentR)
	paymentHandler := handler.NewPaymentHandler(paymentS)

	userR := user.NewRepository(db)
	userS := user.NewService(userR)
	userHandler := handler.NewUserHandler(userS)

	router := gin.Default()

	// router.GET("/", productHandler.Hello)

	roomManager := service.NewRoomManager()
	adapter := adapter.NewGinAdapter(roomManager)

	api := router.Group("/api")
	api.Use(middlewares.JwtAuthMiddleware())
	auth := router.Group("/auth")

	auth.POST("/login", userHandler.Login)
	auth.POST("/register", userHandler.Register)

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
	router.GET("/api/payments/stream", adapter.Stream)

	router.StaticFile("/", "./public/payments.html")
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":3000")

}

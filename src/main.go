package main

import (
	"fmt"
	"log"
	"os"

	"go/src/broadcaster"
	_ "go/src/docs"
	"go/src/handler"
	"go/src/payment"
	"go/src/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 404 custom
func notFound(context *gin.Context) {
	context.JSON(404, gin.H{
		"message": "❌ Page not found ❌",
	})
}

// @title           Product Management API
// @version         1.0
// @description     This is server for products management.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth
func main() {
	// Connection DB
	dbURL := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Println("❌ Connection to DB failed")
		fmt.Println(dbURL)
		log.Fatal(err.Error())
	} else {
		fmt.Println("🔗 Connection to DB OK")
	}

	db.AutoMigrate(&payment.Payment{}, &product.Product{})

	router := gin.Default()

	// Load the pages templates https://gin-gonic.com/docs/examples/html-rendering/
	router.LoadHTMLGlob("src/templates/*")

	// Add the static files
	router.Static("static/", "./src/js")

	// Create the broadcaster
	broadcaster := broadcaster.NewBroadcaster(10)

	// Create the api
	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService, broadcaster)
	api := router.Group("/api")
	{
		payments := api.Group("/payments")
		{
			payments.POST("/", paymentHandler.Create)
			payments.GET("/", paymentHandler.GetAll)
			payments.GET("/sse", paymentHandler.Sse)
			payments.GET("/:id", paymentHandler.GetById)
			payments.PATCH("/:id", paymentHandler.Update)
			payments.DELETE("/:id", paymentHandler.Delete)
		}
		products := api.Group("/products")
		{
			products.POST("/", productHandler.Create)
			products.GET("/", productHandler.GetAll)
			products.GET("/:id", productHandler.GetByID)
			products.PATCH("/:id", productHandler.Update)
			products.DELETE("/:id", productHandler.Delete)
		}
	}

	// Create Web
	web := router.Group("/")
	webHandler := handler.NewWebHandler(productService, paymentService, broadcaster)
	web.GET("/", webHandler.Home)
	web.GET("/products/create", webHandler.CreateProduct)
	web.GET("/payments/create", webHandler.CreatePayment)
	web.GET("/products/:id/edit", webHandler.EditProduct)
	web.GET("/payments/:id/edit", webHandler.EditPayment)

	// Create 404
	router.NoRoute(notFound)

	// Call swagger middleware
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()

	fmt.Println("🚀 Server running on port 3000 🚀")
}

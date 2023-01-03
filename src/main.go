package main

import (
	"fmt"
	"log"
	"os"

	"go/src/broadcaster"
	"go/src/handler"
	"go/src/payment"
	"go/src/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 404 custom
func notFound(context *gin.Context) {
	context.JSON(404, gin.H{
		"message": "❌ Page not found ❌",
	})
}

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
	api := router.Group("/api")
	web := router.Group("/")

	router.NoRoute(notFound)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	api.GET("/products", productHandler.GetAll)
	api.GET("/products/:id", productHandler.GetByID)
	api.POST("/products", productHandler.Create)
	api.PATCH("/products/:id", productHandler.Update)
	api.DELETE("/products/:id", productHandler.Delete)

	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService, broadcaster)
	api.POST("/payments", paymentHandler.Create)
	api.GET("/payments", paymentHandler.GetAll)
	api.GET("/payments/sse", paymentHandler.Sse)
	api.GET("/payments/:id", paymentHandler.GetById)
	api.PATCH("/payments/:id", paymentHandler.Update)
	api.DELETE("/payments/:id", paymentHandler.Delete)

	webHandler := handler.NewWebHandler(productService, paymentService, broadcaster)
	web.GET("/", webHandler.Home)
	web.GET("/products/create", webHandler.CreateProduct)
	web.GET("/payments/create", webHandler.CreatePayment)
	web.GET("/products/:id/edit", webHandler.EditProduct)
	web.GET("/payments/:id/edit", webHandler.EditPayment)

	router.Run()

	fmt.Println("🚀 Server running on port 3000 🚀")
}

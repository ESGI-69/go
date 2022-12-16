package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go/src/handler"
	"go/src/payment"
	"go/src/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// routes
func ping(context *gin.Context) {
	fmt.Println("GET /")
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// 404 custom
func notFound(context *gin.Context) {
	context.JSON(404, gin.H{
		"message": "❌ Page not found ❌",
	})
}

// home page with index.html
func home(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
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

	engine := gin.Default()
	engine.LoadHTMLFiles("src/index.html")

	// Create the api
	api := engine.Group("/api")

	engine.NoRoute(notFound)

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
	paymentHandler := handler.NewPaymentHandler(paymentService)
	api.POST("/payments", paymentHandler.Create)
	api.GET("/payments", paymentHandler.GetAll)
	api.GET("/payments/:id", paymentHandler.GetById)
	api.PATCH("/payments/:id", paymentHandler.Update)
	api.DELETE("/payments/:id", paymentHandler.Delete)

	engine.Run()
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go/src/handler"
	"go/src/payment"

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

// home page
func home(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "🏠 Home 🏠",
	})
}

// 404 custom
func notFound(context *gin.Context) {
	context.JSON(404, gin.H{
		"message": "❌ Page not found ❌",
	})
}

func main() {
	// Connection DB
	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	//migration payment
	db.AutoMigrate(&payment.Payment{})

	//Payment
	paymentRepository := payment.NewRepository(db)
	paymentService := payment.NewService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// Create the gin engine
	engine := gin.Default()
	api := engine.Group("/api")

	engine.NoRoute(notFound)
	engine.GET("/ping", ping)
	engine.GET("/", home)

	api.GET("/payment", paymentHandler.Test)

	engine.Run()
}

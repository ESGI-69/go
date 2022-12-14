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
	engine.LoadHTMLFiles("src/index.html")

	// Create the api
	api := engine.Group("/api")

	// Routes
	engine.NoRoute(notFound)
	engine.GET("/ping", ping)
	engine.GET("/", home)

	// API Routes
	api.GET("/payment", paymentHandler.Test)

	engine.Run()
}

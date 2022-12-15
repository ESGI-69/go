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
	dbURL := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("🔗 Connection to DB OK")
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
	api.POST("/payment", paymentHandler.Create)
	api.GET("/payment", paymentHandler.GetAll)
	api.GET("/payment/:id", paymentHandler.GetById)
	api.PATCH("/payment/:id", paymentHandler.Update)
	api.DELETE("/payment/:id", paymentHandler.Delete)

	engine.Run()
}
